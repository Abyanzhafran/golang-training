package main

import (
	"log"

	handler "golang-assignment/handler/gin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	handler.Handler(router)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to start server: ")
	}
}
