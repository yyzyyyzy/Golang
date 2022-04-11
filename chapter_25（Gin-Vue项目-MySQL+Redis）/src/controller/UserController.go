package controller

import (
	"common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"model"
	"net/http"
	"util"
)

func Register(context *gin.Context) {
	DB := common.ConnectToMySQL()
	defer DB.Close()

	username := context.PostForm("username")
	password := context.PostForm("password")
	phone := context.PostForm("phone")
	if len(username) == 0 {
		util.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		util.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(phone) != 11 {
		util.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	log.Println(username, password, phone)

	//通过手机号判断用户是否已经注册
	if isPhoneExist(DB, phone) {
		util.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	//雪花算法生成唯一user_id
	if err := util.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	user_id, err := util.GetID()
	util.HandleError(err, "util.GetID")

	//密码加密
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		util.Response(context, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	DB.Exec("insert into blog.user(user_id,username,password,phone) values (?,?,?,?);", user_id, username, string(hashedpassword), phone)
	util.Success(context, nil, "注册成功")
}

func isPhoneExist(db *sqlx.DB, phone string) bool {
	var user model.User
	db.Get(&user, "select id, username, password, phone from blog.user where phone = ?;", phone)
	if user.ID != 0 {
		return true
	}
	return false
}

func Login(context *gin.Context) {
	DB := common.ConnectToMySQL()
	defer DB.Close()

	password := context.PostForm("password")
	phone := context.PostForm("phone")
	if len(password) < 6 {
		util.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(phone) != 11 {
		util.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	log.Println(password, phone)

	var user model.User
	DB.Get(&user, "select id, password from blog.user where phone = ?;", phone)
	if user.ID == 0 {
		util.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//表单密码与数据库内的密码进行比对
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		util.Response(context, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token给客户端
	token, err := common.ReleaseToken(user)
	if err != nil {
		util.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	util.Success(context, gin.H{"token": token}, "登录成功")
}

func Info(context *gin.Context) {
	user, _ := context.Get("user")
	util.Success(context, gin.H{"user": util.ToUserDto(user.(model.User))}, "登录成功")
}
