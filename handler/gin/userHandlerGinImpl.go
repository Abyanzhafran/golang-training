package handler

import (
	"net/http"
	"strconv"

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
}

func (handler *UserHandlerImpl) GetAllUsers(ctx *gin.Context) {
	users, err := handler.userUsecase.FindAll(ctx)
	if err != nil {
		// Handle the error and return an Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   users,
	})
}

func (handler *UserHandlerImpl) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	intParam, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	user, err := handler.userUsecase.FindById(ctx, intParam)
	if err != nil {
		// Handle the error and return an Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   user,
	})
}
