package usecase

import (
	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}
