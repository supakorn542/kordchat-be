package services

import (
	"errors"
	"os"
	"time"

	"kordchat-be/models"
	"kordchat-be/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, email string, rawPassword string) error {
	existingUser, _ := repositories.GetUserByEmail(email)

	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	err = repositories.CreateUser(&newUser)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(email string, password string) (string, string, error) {

	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("Invalid email or password")
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv(("JWT_SECRET"))))
	if err != nil {
		return "", "", errors.New("Failed to generate access token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", "", errors.New("Failed to generate refresh token")

	}

	return accessTokenString, refreshTokenString, nil

}

func RefreshToken(refreshTokenString string) (string, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error){
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	if err != nil {
		return  "", errors.New("Invalid refresh token")

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Invalid refresh token")
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": claims["sub"],
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	newAccessTokenString, err := newAccessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return  "", errors.New("Failed to generate new access token")

	}

	return  newAccessTokenString, nil

}



