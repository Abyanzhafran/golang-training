package usecase

import "github.com/gin-gonic/gin"

type submissionUsecase interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}