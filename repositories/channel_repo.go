package repositories

import (
	"kordchat-be/config"
	"kordchat-be/models"

	"github.com/google/uuid"
)

func CreateChannel(channel *models.Channel) error {
	result := config.DB.Create(channel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


func GetChannelsByServerID(serverID uuid.UUID) ([]models.Channel, error) {
	var channels []models.Channel
	result := config.DB.Where("server_id = ?",serverID).Find(&channels)
	if result.Error != nil {
		return  nil, result.Error
	}

	return  channels, nil
}


func GetChannelByID(channelID uuid.UUID) (*models.Channel, error) {
	var channel models.Channel
	result := config.DB.Where("id = ?",channelID).First(&channel)
	if result.Error != nil {
		return nil, result.Error
	}

	return  &channel, nil
}

