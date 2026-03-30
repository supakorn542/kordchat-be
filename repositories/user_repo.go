package repositories

import (
	"kordchat-be/config"
	"kordchat-be/models"
)

func CreateUser(user *models.User) error {
	result := config.DB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}