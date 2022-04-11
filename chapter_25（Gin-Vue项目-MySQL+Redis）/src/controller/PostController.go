package controller

import (
	"common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"model"
	"strconv"
	"util"
)

type MyPostController interface {
	ResultController
}

func NewPostController() MyPostController {
	db := common.ConnectToMySQL()
	return PostController{DB: db}
}

type PostController struct {
	DB *sqlx.DB
}

func (p PostController) Create(context *gin.Context) {
	var requestPost model.Post
	if err := context.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		util.Fail(context, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取登录用户 user
	user, _ := context.Get("user")

	//创建文章
	post := model.Post{
		UserID:     user.(model.User).ID,
		PostID:     requestPost.PostID,
		CategoryID: requestPost.CategoryID,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	//雪花算法生成唯一user_id
	if err := util.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	post_id, err := util.GetID()
	util.HandleError(err, "util.GetID")

	p.DB.Exec("insert into blog.post(user_id,post_id,category_id,title,head_img,content) values (?,?,?,?,?,?);", user.(model.User).ID, post_id, post.CategoryID, post.Title, post.HeadImg, post.Content)
	p.DB.Get(&post, "select user_id,post_id,category_id,title,head_img,content from blog.post where category_id = ?;", user.(model.User).ID, post_id, post.CategoryID, post.Title, post.HeadImg, post.Content, post.CategoryID)
	util.Success(context, gin.H{"user_id": post.UserID, "post_id": post_id, "category_id": post.CategoryID, "title": post.Title, "head_img": post.HeadImg, "content": post.Content}, "创建成功")
}

func (p PostController) Update(context *gin.Context) {
	var requestPost model.Post
	if err := context.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		util.Fail(context, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path的参数
	postID, _ := strconv.Atoi(context.Params.ByName("id"))
	var updatePost model.Post
	err := p.DB.Get(&updatePost, "select post_id, user_id, category_id, title, head_img, content from blog.post where id = ?;", postID)
	if err == nil {
		util.Fail(context, nil, "文章不存在")
		return
	}

	////判断当前用户是否为文章的作者
	//user, _ := context.Get("user") //获取登录用户 user
	//userID := user.(model.User).ID
	//if userID != updatePost.UserID {
	//	util.Fail(context, nil, "文章不属于您，请勿非法操作")
	//	return
	//}

	//更新文章
	p.DB.Exec("update blog.post set category_id = ? where id = ?;", requestPost.CategoryID, postID)
	p.DB.Exec("update blog.post set title = ? where id = ?;", requestPost.Title, postID)
	p.DB.Exec("update blog.post set head_img = ? where id = ?;", requestPost.HeadImg, postID)
	p.DB.Exec("update blog.post set content = ? where id = ?;", requestPost.Content, postID)

	util.Success(context, gin.H{"category_id": requestPost.CategoryID, "title": requestPost.Title, "head_img": requestPost.HeadImg, "content": requestPost.Content}, "更新成功")
}

func (p PostController) Query(context *gin.Context) {
	//获取path的参数
	postID, _ := strconv.Atoi(context.Params.ByName("id"))
	var post model.Postman
	if err := p.DB.Get(&post, "select id from blog.post where id = ?;", postID); err != nil {
		util.Fail(context, nil, "文章不存在")
		return
	}
	util.Success(context, gin.H{"post": post}, "查看成功")
}

func (p PostController) Delete(context *gin.Context) {
	//获取path的参数
	postID := context.Params.ByName("id")
	var post model.Post
	if err := p.DB.Get(&post, "select id from blog.post where id = ?;", postID); err != nil {
		util.Fail(context, nil, "文章不存在")
		return
	}
	//判断当前用户是否为文章的作者
	user, _ := context.Get("user") //获取登录用户 user
	userID := user.(model.User).ID
	if userID != post.UserID {
		util.Fail(context, nil, "文章不属于您，请勿非法操作")
		return
	}

	p.DB.Exec("delete from blog.post where id = ?;", postID)

	util.Success(context, gin.H{"post": post}, "删除成功")
}
