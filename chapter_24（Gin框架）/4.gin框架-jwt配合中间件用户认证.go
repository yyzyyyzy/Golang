package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 指定加密密钥
var jwtKey = []byte("a_secret_key")

// Claims 是用户的状态和额外的元数据（官方字段）
// 如果需要额外记录别的信息，可以添加到这个结构体中
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user User) (string, error) {
	expirationtime := time.Now().Add(7 * 24 * time.Hour) //过期时间
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationtime.Unix(),
			//token发行日期
			IssuedAt: time.Now().Unix(),
			//token发行人
			Issuer: "李子康编程大魔王",
			//token的发行的主题
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//token.SignedString()内部生成签名字符串，再用于获取完整、已签名的token
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

// 中间件配合token验证
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 1.获取authorization header
		tokenString := context.GetHeader("Authorization")

		// 2.验证token格式（如果1.tokenString为空 2.不是以Bearer（请求头）开头那么就判断没有传token，或者是以错误的格式开头）
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort() //抛弃此次请求
			return
		}

		tokenString = tokenString[7:] //"Bearer "7个字节
		token, claims, err := ParseToken(tokenString)
		// 如果1.解析不到token 2.token格式有误 那么九抛弃此次请求
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
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
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		context.Set("user", user)
		context.Next()
	}
}
