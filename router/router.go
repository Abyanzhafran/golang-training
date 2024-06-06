package router

import (
	"example/hello/handler"
	"example/hello/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	privateApi := r.Group("/private")
	privateApi.Use(middleware.AuthMiddleware())
	{
		privateApi.POST("/post", handler.PostHandler)
	}
}
