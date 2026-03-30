package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"kordchat-be/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("Warning: Could not load .env.local file, using system environment variables instead")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected successfully")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Server{},
		&models.Channel{},
		&models.Message{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migrated successfully")

}
