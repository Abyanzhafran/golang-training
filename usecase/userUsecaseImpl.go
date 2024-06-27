package usecase

import (
	"fmt"
	"net/http"
	"strconv"

	"golang-assignment/entity"
	"golang-assignment/repository"

	"github.com/gin-gonic/gin"
)

type UserUsecaseImpl struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(UserRepo repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{UserRepo: UserRepo}
}

func (usecase *UserUsecaseImpl) FindAll(ctx *gin.Context) {
	users, err := usecase.UserRepo.GetAll(ctx)
	if err != nil {
		// Handle the error and return an Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   users,
	})
}

func (usecase *UserUsecaseImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	// Converting the string parameter to int64
	intParam, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
		return
	}

	user, err := usecase.UserRepo.GetById(ctx, intParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (usecase *UserUsecaseImpl) Create(ctx *gin.Context) {
	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := usecase.UserRepo.Create(ctx, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   user,
	})
}

func (usecase *UserUsecaseImpl) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	// Converting the string parameter to int64
	intParam, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
		return
	}

	user, err := usecase.UserRepo.GetById(ctx, intParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "User not found",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := usecase.UserRepo.Update(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   user,
	})
}

func (usecase *UserUsecaseImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	// Converting the string parameter to int64
	intParam, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
		return
	}

	user, err := usecase.UserRepo.GetById(ctx, intParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := usecase.UserRepo.Delete(ctx, intParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"error":  "",
		"data":   user,
	})
}
