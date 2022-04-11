package controller

import "github.com/gin-gonic/gin"

type ResultController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Query(context *gin.Context)
	Delete(context *gin.Context)
}
