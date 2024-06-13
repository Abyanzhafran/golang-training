package main

import (
	"log"

	"golang-advance/entity"
	"golang-advance/handler"
	"golang-advance/middleware"
	"golang-advance/repository"
	"golang-advance/router"
	"golang-advance/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	var mockUserDBInSlice []entity.User
	userRepo := repository.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Running On Port 8080")

	r.Run(":8080")
}
