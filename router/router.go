package router

import (
	"golang-advance/handler"
	"golang-advance/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)

	usersPrivateEndpoint := r.Group("/users")

	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	usersPrivateEndpoint.POST("", userHandler.CreateUser)
	usersPrivateEndpoint.POST("/", userHandler.CreateUser)
	usersPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
	usersPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)
}
