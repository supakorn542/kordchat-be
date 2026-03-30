package main

import (
	"log"

	"kordchat-be/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	log.Println("Server is running on port 8080...")
	r.Run(":8080")
}
