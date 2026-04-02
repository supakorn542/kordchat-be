package controllers

import (
	"kordchat-be/dtos"
	"kordchat-be/services"

	"github.com/gin-gonic/gin"

	"net/http"
)

// CreateServer godoc
// @Summary 	Create new server
// @Description request server name to create new server
// @Tags 		Server
// @Accept 		json
// @Produce 	json
// @Param		input body dtos.CreateServerRequest true "CreateServer"
// @Success 	201 {object} map[string]interface{} "create server successfully"
// @Failure 	400 {object} map[string]interface{} "invalid data"
// @Failure 	401 {object} map[string]interface{} "unauthorized"
// @Router 		/servers [post]
func CreateServer(c *gin.Context) {
	var input dtos.CreateServerRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userIDStr := userID.(string)

	server, err := services.CreateServer(input.Name, userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "create server successfully",
		"server":  server,
	})
}

// GetServersByUserID godoc
// @Summary    		Get User's Servers
// @Description 	Get servers list by userID
// @Tags 			Server
// @Accept 			json
// @Produce 		json
// @Success      	200  {array} dtos.ServerResponse "Servers List"
// @Failure 		400 {object} map[string]interface{} "invalid data"
// @Failure      	401  {object} map[string]interface{} "unauthorized"
// @Router       	/servers [get]
func GetServersByUserID(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userIDStr := userID.(string)

	servers, err := services.GetServersByUserID(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	response := dtos.ToServerResponses(servers)

	c.JSON(http.StatusOK, response)
}
