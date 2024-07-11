package usecase

import (
	"context"
	"fmt"

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

func (usecase *UserUsecaseImpl) FindAll(ctx *gin.Context) ([]entity.User, error) {
	users, err := usecase.UserRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all users: %v", err)
	}
	return users, nil
}

func (usecase *UserUsecaseImpl) FindById(ctx *gin.Context, id int) (entity.User, error) {
	// id := ctx.Param("id")

	// convert string to int
	// intParam, err := strconv.Atoi(id)
	// if err != nil {
	// 	return entity.User{}, fmt.Errorf("failed to convert string to int : %v", err)
	// }

	user, err := usecase.UserRepo.GetById(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to fetch all users: %v", err)
	}
	return user, nil
}

func (usecase *UserUsecaseImpl) Create(ctx context.Context, user *entity.User) (entity.User, error) {
	createdUser, err := usecase.UserRepo.Create(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal membuat pengguna: %v", err)
	}
	return createdUser, nil
}

// func (usecase *UserUsecaseImpl) Update(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	// Converting the string parameter to int
// 	intParam, err := strconv.Atoi(id)
// 	if err != nil {
// 		fmt.Println("Error converting string to int64:", err)
// 		return
// 	}

// 	user, err := usecase.UserRepo.GetById(ctx, intParam)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"status": false,
// 			"error":  "User not found",
// 		})
// 		return
// 	}

// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"status": false,
// 			"error":  err.Error(),
// 		})
// 		return
// 	}

// 	if err := usecase.UserRepo.Update(ctx, user); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"status": false,
// 			"error":  err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"status": true,
// 		"error":  "",
// 		"data":   user,
// 	})
// }

// func (usecase *UserUsecaseImpl) Delete(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	// Converting the string parameter to int
// 	intParam, err := strconv.Atoi(id)
// 	if err != nil {
// 		fmt.Println("Error converting string to int64:", err)
// 		return
// 	}

// 	user, err := usecase.UserRepo.GetById(ctx, intParam)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"status": false,
// 			"error":  err.Error(),
// 		})
// 		return
// 	}

// 	if err := usecase.UserRepo.Delete(ctx, intParam); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"status": false,
// 			"error":  err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"status": true,
// 		"error":  "",
// 		"data":   user,
// 	})
// }
