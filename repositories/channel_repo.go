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


func GetChannelsByServerID(serverID string) ([]models.Channel, error) {
	var channels []models.Channel
	result := config.DB.Where("server_id = ?",serverID).Find(&channels)
	if result.Error != nil {
		return  nil, result.Error
	}

	return  channels, nil
}

