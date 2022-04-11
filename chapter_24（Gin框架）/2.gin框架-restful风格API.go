package main

import "github.com/gin-gonic/gin"

//1.获取文章 /blog/getXxx Get blog/Xxx
//2.添加 /blog/addXxx POST blog/Xxx
//3.修改 /blog/updateXxx PUT blog/Xxx
//4.删除 /blog/delXxxx DELETE blog/Xxx

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.Run(":8080")
}
