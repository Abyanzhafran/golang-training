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

	submissionRepo := repository.NewSubmissionRepository(db)
	submissionUsecase := usecase.NewSubmissionUsecase(submissionRepo, userRepo)

	userHandler := router.Group("/users")
	userHandler.GET("", userUsecase.FindAll)
	userHandler.GET("/:id", userUsecase.FindById)
	userHandler.POST("", userUsecase.Create)
	userHandler.PUT("/:id", userUsecase.Update)
	userHandler.DELETE("/:id", userUsecase.Delete)

	submissionRepository := router.Group("/submissions")
	submissionRepository.GET("", submissionUsecase.FindAll)
	submissionRepository.POST("", submissionUsecase.Create)
}
