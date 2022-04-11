package controller

import (
	"common"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"model"
	"strconv"
	"util"
)

type MyCategoryController interface {
	ResultController
}

type CategoryController struct { //结构体实现MyCategoryController接口
	DB *sqlx.DB
}

func NewCategoryController() MyCategoryController {
	db := common.ConnectToMySQL()
	return CategoryController{DB: db}
}

func (c CategoryController) Create(context *gin.Context) {
	var requestCategory model.Category
	context.ShouldBind(&requestCategory)

	if requestCategory.Name == "" {
		util.Fail(context, nil, "数据验证错误，分类名称必填")
		return
	}
	c.DB.Exec("insert into blog.category (name) values (?);", requestCategory.Name)
	c.DB.Get(&requestCategory, "select id,name from blog.category where name = ?;", requestCategory.Name)
	util.Success(context, gin.H{"category": gin.H{"ID": requestCategory.ID, "Name": requestCategory.Name}}, "")

}

func (c CategoryController) Update(context *gin.Context) {
	//绑定body 中的参数
	var requestCategory model.Category
	context.ShouldBind(&requestCategory)

	if requestCategory.Name == "" {
		util.Fail(context, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path 中的参数
	categoryID, _ := strconv.Atoi(context.Params.ByName("id"))
	var updateCategory model.Category
	err := c.DB.Get(&updateCategory, "select id,name from blog.category where id = ?;", categoryID)
	if err != nil {
		util.Fail(context, nil, "分类不存在")
		return
	}
	//更新分类
	c.DB.Exec("update blog.category set name = ? where id = ?;", requestCategory.Name, categoryID)
	util.Success(context, gin.H{"ID": categoryID, "Name": requestCategory.Name}, "修改成功")
}

func (c CategoryController) Query(context *gin.Context) {
	// 获取path的参数

	categoryID, _ := strconv.Atoi(context.Params.ByName("id"))

	var category model.Category
	err := c.DB.Get(&category, "select id,name from blog.category where id = ?;", categoryID)
	if err != nil {
		util.Fail(context, nil, "分类不存在")
		return
	}

	//更新分类
	util.Success(context, gin.H{"ID": category.ID, "Name": category.Name}, "")

}

func (c CategoryController) Delete(context *gin.Context) {
	categoryID, _ := strconv.Atoi(context.Params.ByName("id"))

	var category model.Category
	if err := c.DB.Get(&category, "select id,name from blog.category where id = ?;", categoryID); err != nil {
		util.Fail(context, nil, "分类不存在")
		return
	}

	c.DB.Exec("delete from blog.category where id = ?;", categoryID)
	util.Success(context, nil, "删除成功")

}
