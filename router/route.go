package route

import (
	"golang-assignment/config"
	"golang-assignment/repository"
	"golang-assignment/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	db := config.NewDB()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	userHandler := router.Group("/users")
	userHandler.GET("", userUsecase.FindAll)
	userHandler.GET("/:id", userUsecase.FindById)
	userHandler.POST("", userUsecase.Create)
	userHandler.PUT("/:id", userUsecase.Update)
	userHandler.DELETE("/:id", userUsecase.Delete)
}
