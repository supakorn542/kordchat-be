package services

import (
	"kordchat-be/models"
	"kordchat-be/repositories"

	"github.com/google/uuid"
)

func CreateServer(name string, ownerIDStr string) (*models.Server, error) {
	ownerID, err := uuid.Parse(ownerIDStr)
	if err != nil {
		return nil, err
	}

	owner, err := repositories.GetUserByID(ownerIDStr)
	if err != nil {
		return nil, err
	}

	newServer := models.Server{
		Name:    name,
		OwnerID: ownerID,
		Users:   []models.User{*owner},
	}

	err = repositories.CreateServer(&newServer)
	if err != nil {
		return nil, err
	}

	return &newServer, nil

}

func GetServersByUserID(userIDStr string) ([]models.Server, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	result, err := repositories.GetServersByUserID(userID)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func AddUserToServer(serverIDStr string, userIDStr string) (*models.Server, error) {
	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	server, err := repositories.GetServerByID(serverID)
	if err != nil {
		return nil, err
	}

	err = repositories.AddUserToServer(serverID, userID)
	if err != nil {
		return nil, err
	}

	return server, nil

}
