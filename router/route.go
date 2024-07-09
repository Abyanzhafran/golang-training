package route

import (
	"golang-assignment/config"
	handler "golang-assignment/handler/gin"
	"golang-assignment/repository"
	"golang-assignment/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	db := config.NewDB()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	route := router.Group("/users")
	route.GET("", userHandler.GetAllUsers)
	route.GET("/:id", userUsecase.FindById)
	route.POST("", userHandler.CreateUser)
	route.PUT("/:id", userUsecase.Update)
	route.DELETE("/:id", userUsecase.Delete)
}
