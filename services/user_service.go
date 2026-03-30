package services

import (
	"errors"

	"kordchat-be/models"
	"kordchat-be/repositories"

	"golang.org/x/crypto/bcrypt"

)


func RegisterUser(username string, email string, rawPassword string) error {
	existingUser, _ := repositories.GetUserByEmail(email)

	if(existingUser != nil){
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Username: username,
		Password: string(hashedPassword),
		Email: email,
	}

	err = repositories.CreateUser(&newUser)
	if(err != nil){
		return err
	}

	return nil
}