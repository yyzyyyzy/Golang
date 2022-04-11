package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	db := ConnToMySQL()
	defer db.Close()

	r := gin.Default()
	//注册路由
	r.POST("/register", func(context *gin.Context) {

		//1.获取表单参数
		name := context.PostForm("name")
		password := context.PostForm("password")
		phone := context.PostForm("phone")

		//2.表单参数验证
		if len(phone) != 11 {
			response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
			return
		}
		if len(password) < 6 {
			response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
			return
		}
		//名称没有传，就提供10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, password, phone)

		//3.判断手机号是否存在：1.不存在则增加一条新数据 2.存在返回422
		if isPhoneExist(db, phone) {
			response(context, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
			return
		}

		//4.用户密码加密
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			response(context, http.StatusInternalServerError, 500, nil, "加密错误")
			return
		}

		db.Exec("insert into user(name,password,phone) values (?,?,?);", name, string(hashedpassword), phone)
		Success(context, nil, "注册成功")
	})
	//登录路由
	r.POST("/login", func(context *gin.Context) {

		//1.获取参数
		password := context.PostForm("password")
		phone := context.PostForm("phone")

		//2.表单验证
		if len(phone) != 11 {
			response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
			return
		}
		if len(password) < 6 {
			response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
			return
		}

		//3.判断手机号是否存在：1.存在则登陆成功 2.不存在则返回不存在
		var user User
		db.Get(&user, "SELECT id, name, password, phone FROM user where Phone = ?;", phone)
		if user.ID == 0 {
			response(context, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
			return
		}

		//4.判断密码是否正确：1.密码正确发放token给前端 2.密码错误返回密码错误
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			response(context, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		}

		//5.发放token
		token, err := ReleaseToken(user)
		if err != nil {
			response(context, http.StatusInternalServerError, 500, nil, "系统异常")
			log.Printf("token generate error : %v", err)
			return
		}

		//6.返回结果
		Success(context, gin.H{"token": token}, "登陆成功")
	})

	//token验证路由
	r.GET("/info", AuthMiddleware(), func(context *gin.Context) {
		// 从上下文中获取用户信息
		user, _ := context.Get("user")
		Success(context, gin.H{"user": ToUserDto(user.(User))}, "")
		//context.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": ToUserDto(user.(User))}})
	})

	panic(r.Run(":8080"))
}

type User struct {
	ID       uint   `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
}

func ConnToMySQL() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/cbb")
	HandleSQLError(err, "sqlx.Open")
	return db
}

func HandleSQLError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func isPhoneExist(db *sqlx.DB, phone string) bool {
	var user User
	db.Get(&user, "SELECT id, name, password, phone FROM user where Phone = ?;", phone)
	if user.ID != 0 {
		return true
	}
	return false
}

//jwt用户认证
var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user User) (string, error) {
	expirationtime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "李子康编程大魔王",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 1.获取authorization header
		tokenString := context.GetHeader("Authorization")

		// 2.验证token格式（如果1.tokenString为空 2.不是以Bearer（请求头）开头那么就判断没有传token，或者是以错误的格式开头）
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response(context, http.StatusUnauthorized, 401, nil, "权限不足")
			context.Abort() //抛弃此次请求
			return
		}

		tokenString = tokenString[7:] //"Bearer "7个字节
		token, claims, err := ParseToken(tokenString)
		// 如果1.解析不到token 2.token格式有误 那么九抛弃此次请求
		if err != nil || !token.Valid {
			response(context, http.StatusUnauthorized, 401, nil, "权限不足")
			context.Abort()
			return
		}

		//如果token通过验证，获取claim中的userId
		userId := claims.UserId
		DB := ConnToMySQL()
		var user User
		DB.Get(&user, "SELECT id, name, password, phone FROM user where id = ?;", userId)

		//验证用户是否存在：如果用户不存在，token无效 否则用户存在，将user信息写入上下文
		if userId == 0 {
			response(context, http.StatusUnauthorized, 401, nil, "权限不足")
			context.Abort()
			return
		}
		context.Set("user", user)
		context.Next()
	}
}

// RandomString 生成随机10位的字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// UserDto User结构体转为UserDto结构体
type UserDto struct {
	Name  string `db:"name"`
	Phone string `db:"phone"`
}

func ToUserDto(user User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}

func response(context *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	context.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(context *gin.Context, data gin.H, msg string) {
	response(context, http.StatusOK, 200, data, msg)
}

func Fail(context *gin.Context, data gin.H, msg string) {
	response(context, http.StatusOK, 400, data, msg)
}
