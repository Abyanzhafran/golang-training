package main

import (
	"log"

	"example/hello/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.SetupRouter(r)

	log.Println("Running server on port 8080")
	r.Run(":8080")
}
