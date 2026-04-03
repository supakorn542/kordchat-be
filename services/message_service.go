package services

import (
	"kordchat-be/models"
	"kordchat-be/repositories"

	"github.com/google/uuid"

	"errors"
)

func CreateMessage(content string, channelIDStr string, userIDStr string) (*models.Message, error) {
	channelID, err := uuid.Parse(channelIDStr)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	channel, err := repositories.GetChannelByID(channelID)
	if err != nil {
		return nil, err
	}

	if channel.Type != "text" {
		return nil, errors.New("cannot send messages to a non-text channel")
	}

	userServers, err := repositories.GetServersByUserID(userID)
	if err != nil {
		return nil, err
	}

	isMember := false
	for _, server := range userServers {
		if server.ID == channel.ServerID {
			isMember = true
			break
		}
	}

	if !isMember {
		return nil, errors.New("user is not a member of the server")
	}

	newMessage := models.Message{
		Content:   content,
		ChannelID: channelID,
		UserID:    userID,
	}

	err = repositories.CreateMessage(&newMessage)
	if err != nil {
		return nil, err
	}

	user, err := repositories.GetUserByID(userIDStr)
	if err == nil {
		newMessage.User = *user
	}

	return &newMessage, nil

}

func GetMessagesByChannelID(channelIDStr string, userIDStr string) ([]models.Message, error) {
	channelID, err := uuid.Parse(channelIDStr)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	channel, err := repositories.GetChannelByID(channelID)
	if err != nil {
		return nil, err
	}

	userServers, err := repositories.GetServersByUserID(userID)
	if err != nil {
		return nil, err
	}

	isMember := false
	for _, server := range userServers {
		if server.ID == channel.ServerID {
			isMember = true
			break
		}
	}

	if !isMember {
		return nil, errors.New("user is not a member of the server")
	}

	messages, err := repositories.GetMessagesByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
