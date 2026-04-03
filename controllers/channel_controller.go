package controllers

import (
	"kordchat-be/dtos"
	"kordchat-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateChannel godoc
// @Summary 		Create a new channel
// @Description 	Create a new channel in a server
// @Tags 			Channel
// @Accept 			json
// @Produce 		json
// @Param        	serverId path string true "Server ID"
// @Param 			input body dtos.CreateChannelRequest true "CreateChannel"
// @Success 		201 {object} dtos.ChannelResponse "Channel created successfully"
// @Failure			400 {object} map[string]interface{} "Invalid data"
// @Failure 		401 {object} map[string]interface{} "Unauthorized"
// @Router /servers/{serverId}/channels [post]
func CreateChannel(c *gin.Context) {
	var input dtos.CreateChannelRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serverIDStr := c.Param("serverId")

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userIDStr := userID.(string)

	channel, err := services.CreateChannel(input.Name, input.Type, serverIDStr, userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dtos.ToChannelResponse(*channel)

	c.JSON(http.StatusCreated, response)

}

// GetChannelsByServerID godoc
// @Summary 		Get channels by server ID
// @Description 	Get all channels in a server
// @Tags 			Channel
// @Accept 			json
// @Produce 		json
// @Param        	serverId path string true "Server ID"
// @Success 		200 {array} dtos.ChannelResponse "Channels retrieved successfully"
// @Failure			400 {object} map[string]interface{} "Invalid data"
// @Failure 		401 {object} map[string]interface{} "Unauthorized"
// @Router /servers/{serverId}/channels [get]
func GetChannelsByServerID(c *gin.Context) {
	serverIDStr := c.Param("serverId")

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDStr := userID.(string)

	channels, err := services.GetChannelsByServerID(serverIDStr, userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dtos.ToChannelResponses(channels)

	c.JSON(http.StatusOK, response)

}
