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
