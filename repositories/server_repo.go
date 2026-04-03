package repositories

import (
	"kordchat-be/config"
	"kordchat-be/models"

	"github.com/google/uuid"
)

func CreateServer(server *models.Server) error {
	result := config.DB.Create(server)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetServersByUserID(userID uuid.UUID) ([]models.Server, error) {
	var servers []models.Server

	err := config.DB.Model(&models.User{ID: userID}).Association("Servers").Find(&servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}

func AddUserToServer(serverID uuid.UUID, userID uuid.UUID) error {
	var server models.Server
	var user models.User

	if err := config.DB.First(&server, "id = ?", serverID).Error; err != nil {
		return err
	}

	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		return err
	}

	err := config.DB.Model(&server).Association("Users").Append(&user)
	return err
}

func GetServerByID(serverID uuid.UUID) (*models.Server, error) {
	var server models.Server
	result := config.DB.First(&server, "id = ?", serverID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &server, nil
}
