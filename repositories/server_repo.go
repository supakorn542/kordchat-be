package repositories

import (
	"kordchat-be/config"
	"kordchat-be/models"

	"github.com/google/uuid"

)

func CreateServer(server *models.Server) error {
	result := config.DB.Create(server)
	if result.Error != nil {
		return  result.Error
	}

	return nil
}

func GetServersByUserID(userID uuid.UUID) ([]models.Server, error){
	var servers []models.Server

	err := config.DB.Model(&models.User{ID: userID}).Association("Servers").Find(&servers)
	if err != nil {
		return  nil, err
	}

	return  servers, nil
}