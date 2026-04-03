package services

import (
	"errors"
	"kordchat-be/models"
	"kordchat-be/repositories"

	"github.com/google/uuid"
)

func CreateChannel(name string, channelType string, serverIDStr string, userIDStr string) (*models.Channel, error) {
	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	userServers, err := repositories.GetServersByUserID(userID)
	if err != nil {
		return nil, err
	}

	isMember := false
	for _, server := range userServers {
		if server.ID == serverID {
			isMember = true
			break
		}
	}

	if !isMember {
		return nil, errors.New("user is not a member of the server")
	}

	newChannel := models.Channel{
		Name:     name,
		Type:     channelType,
		ServerID: serverID,
	}

	err = repositories.CreateChannel(&newChannel)
	if err != nil {
		return nil, err
	}

	return &newChannel, nil

}

func GetChannelsByServerID(serverIDStr string, userIDStr string) ([]models.Channel, error) {

	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	userServers, err := repositories.GetServersByUserID(userID)
	if err != nil {
		return nil, err
	}

	isMember := false
	for _, server := range userServers {
		if server.ID == serverID {
			isMember = true
			break
		}
	}

	if !isMember {
		return nil, errors.New("user is not a member of the server")
	}

	channels, err := repositories.GetChannelsByServerID(serverID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
