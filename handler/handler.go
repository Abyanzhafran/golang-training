package handler

import (
	"golang-assignment/config"
	"golang-assignment/repository"
	"golang-assignment/usecase"

	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	db := config.NewDB()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	userHandler := router.Group("/users")
	userHandler.GET("", userUsecase.FindAll)
	userHandler.POST("", userUsecase.Create)
}
