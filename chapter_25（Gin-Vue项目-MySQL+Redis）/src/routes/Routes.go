package routes

import (
	"controller"
	"github.com/gin-gonic/gin"
	"middleware"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("user")
	{
		userGroup.POST("/register", controller.Register)
		userGroup.POST("/login", controller.Login)
		userGroup.GET("/info", middleware.AuthMiddleware(), controller.Info)
	}

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("/add", categoryController.Create)
	categoryRoutes.DELETE("/delete/:id", categoryController.Delete)
	categoryRoutes.PUT("/update/:id", categoryController.Update)
	categoryRoutes.GET("/show/:id", categoryController.Query)

	postRoutes := r.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware()) //Use方法使用中间件
	postController := controller.NewPostController()
	postRoutes.POST("/add", postController.Create)
	postRoutes.DELETE("/delete/:id", postController.Delete)
	postRoutes.PUT("/update/:id", postController.Update)
	postRoutes.GET("/show/:id", postController.Query)
	return r
}
