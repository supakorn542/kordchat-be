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
		return  nil, err
	}

	return  &newServer, nil

}
