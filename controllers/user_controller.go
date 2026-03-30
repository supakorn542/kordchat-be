package controllers

import (
	"github.com/gin-gonic/gin"
	"kordchat-be/dtos"
	"kordchat-be/services"
	"net/http"
)

// Register godoc
// @Summary      User register
// @Description  request Username, Email, Password to create User
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input body dtos.RegisterRequest true "Register"
// @Success      201  {object}  map[string]interface{} "register user successfully"
// @Failure      400  {object}  map[string]interface{} "invalid data"
// @Router       /register [post]
func Register(c *gin.Context) {
	var input dtos.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.RegisterUser(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register user successfully",
	})
}


// Login godoc
// @Summary      Login
// @Description  validate email and password to create JWT Token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input body dtos.LoginRequest true "Login data"
// @Success      200  {object}  map[string]interface{} "User login successfully"
// @Failure      400  {object}  map[string]interface{} "Invalid Data"
// @Failure      401  {object}  map[string]interface{} "Invalid email or password"
// @Router       /login [post]
func Login(c *gin.Context){
	var input dtos.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := services.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"access_token": accessToken,
		"refresh_token": refreshToken,
	})
}


// RefreshToken godoc
// @Summary      create new access token
// @Description  use refresh token to create new access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input body dtos.RefreshRequest true "Refresh Token"
// @Success      200  {object}  map[string]interface{} "create access token successfully"
// @Failure      401  {object}  map[string]interface{} "invalid refresh token"
// @Router       /refresh [post]
func RefreshToken(c *gin.Context){
	var input dtos.RefreshRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"please send refresh token"})
		return
	}

	newAccessToken, err := services.RefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
