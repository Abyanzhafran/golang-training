package main

import (
	"log"

	route "golang-assignment/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()

	ginRouter.Use(cors.Default())

	route.SetupRouter(ginRouter)

	if err := ginRouter.Run(":3000"); err != nil {
		log.Fatalf("failed to start server: ")
	}
}
