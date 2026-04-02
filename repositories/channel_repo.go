package repositories

import (
	"kordchat-be/config"
	"kordchat-be/models"
)

func CreateChannel(channel *models.Channel) error {
	result := config.DB.Create(channel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
