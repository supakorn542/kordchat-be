package repositories

import (
	"kordchat-be/config"
	"kordchat-be/models"

)

func CreateServer(server *models.Server) error {
	result := config.DB.Create(server)
	if result.Error != nil {
		return  result.Error
	}

	return nil
}