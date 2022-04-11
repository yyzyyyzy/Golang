package middleware

import (
	"common"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"strings"
	"util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 1.获取authorization header
		tokenString := context.GetHeader("Authorization")

		// 2.验证token格式（如果1.tokenString为空 2.不是以Bearer（请求头）开头那么就判断没有传token，或者是以错误的格式开头）
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			util.Response(context, http.StatusUnauthorized, 401, nil, "权限不足1")
			context.Abort() //抛弃此次请求
			return
		}

		tokenString = tokenString[7:] //"Bearer "7个字节
		token, claims, err := common.ParseToken(tokenString)
		// 如果1.解析不到token 2.token格式有误 那么就抛弃此次请求
		if err != nil || !token.Valid {
			util.Response(context, http.StatusUnauthorized, 401, nil, "权限不足2")
			context.Abort()
			return
		}

		//如果token通过验证，获取claim中的userId
		userId := claims.UserId
		DB := common.ConnectToMySQL()
		var user model.User
		DB.Get(&user, "SELECT id, username, password, phone FROM user where id = ?;", userId)

		//验证用户是否存在：如果用户不存在，token无效 否则用户存在，将user信息写入上下文
		if userId == 0 {
			util.Response(context, http.StatusUnauthorized, 401, nil, "权限不足3")
			context.Abort()
			return
		}
		context.Set("user", user)
		context.Next()
	}
}
