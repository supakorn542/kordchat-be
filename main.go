package main

import (
	"log"

	"kordchat-be/config"
	"kordchat-be/routes"

	_ "kordchat-be/docs"

	"github.com/gin-gonic/gin"
)

// @title           KordChat API
// @version         1.0
// @description     API Server KordChat
// @host            localhost:8080
// @BasePath        /api
func main() {

	config.ConnectDB()

	r := gin.Default()


	routes.SetupRoutes(r)

	log.Println("Server is running on port 8080...")
	r.Run(":8080")
}
