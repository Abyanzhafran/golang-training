package handler

import (
	"net/http"

	"golang-assignment/entity"
	"golang-assignment/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandlerImpl {
	return UserHandlerImpl{userUsecase: userUsecase}
}

// CreateUser menghandle permintaan untuk membuat user baru
func (handler *UserHandlerImpl) CreateUser(ctx *gin.Context) {
	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	createdUser, err := handler.userUsecase.Create(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdUser)

	return
}
