package router

import (
	"example/hello/handler"
	"example/hello/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	publicApi := r.Group("/users")
	publicApi.GET("/:id", handler.GetUser)
	publicApi.GET("/", handler.GetAllUsers)

	privateApi := r.Group("/users")
	privateApi.Use(middleware.AuthMiddleware())
	privateApi.POST("/", handler.CreateUser)
	privateApi.PUT("/:id", handler.UpdateUser)
	privateApi.DELETE("/:id", handler.DeleteUser)
}
