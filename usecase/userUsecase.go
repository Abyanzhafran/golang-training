package usecase

import (
	"context"

	"golang-assignment/entity"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Create(ctx context.Context, user *entity.User) (entity.User, error)
	FindAll(ctx *gin.Context) ([]entity.User, error)
	FindById(ctx *gin.Context, id int) (entity.User, error)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}
