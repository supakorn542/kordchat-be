package repositories

import (
	"kordchat-be/config"
	"kordchat-be/models"

	"github.com/google/uuid"
)

func CreateMessage(message *models.Message) error {
	result := config.DB.Create(message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


func GetMessagesByChannelID(channelID uuid.UUID) ([]models.Message, error){
	var messages []models.Message
	result := config.DB.Preload("User").Where("channel_id = ?", channelID).Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

