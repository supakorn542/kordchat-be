package main

import (
	"log"

	"kordchat-be/config"
	"kordchat-be/routes"

	_ "kordchat-be/docs"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// @title           KordChat API
// @version         1.0
// @description     API Server KordChat
// @host            localhost:8080
// @BasePath        /api
func main() {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Println("Server is running on port 8080...")
	r.Run(":8080")
}
