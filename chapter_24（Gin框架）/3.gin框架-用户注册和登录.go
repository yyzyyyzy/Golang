package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"os"
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
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		//名称没有传，就提供10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, password, phone)

		//3.判断手机号是否存在：1.不存在则增加一条新数据 2.存在返回422
		if isPhoneExist(db, phone) {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		//4.用户密码加密
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
			return
		}

		db.Exec("insert into user(name,password,phone) values (?,?,?);", name, string(hashedpassword), phone)
		context.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
	})
	//登录路由
	r.POST("/login", func(context *gin.Context) {

		//1.获取参数
		password := context.PostForm("password")
		phone := context.PostForm("phone")

		//2.表单验证
		if len(phone) != 11 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		//3.判断手机号是否存在：1.存在则登陆成功 2.不存在则返回不存在
		var user User
		db.Get(&user, "SELECT id, name, password, phone FROM user where Phone = ?;", phone)
		if user.ID == 0 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
			return
		}

		//4.判断密码是否正确：1.密码正确发放token给前端 2.密码错误返回密码错误
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		}

		//5.发放token
		token := "666"

		//6.返回结果
		context.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登陆成功"})
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
