package controllers

import (
	"net/http"

	"kordchat-be/dtos"
	"kordchat-be/services"
	"kordchat-be/websockets"

	"github.com/gin-gonic/gin"
)

// CreateMessage godoc
// @Summary 		Create a new message
// @Description 	Create a new message in a channel
// @Tags 			Message
// @Accept 			json
// @Produce 		json
// @Param        	channelId path string true "Channel ID"
// @Param 			input body dtos.CreateMessageRequest true "CreateMessage"
// @Success 		201 {object} dtos.MessageResponse "Message created successfully"
// @Failure			400 {object} map[string]interface{} "Invalid data"
// @Failure 		401 {object} map[string]interface{} "Unauthorized"
// @Router /channels/{channelId}/messages [post]
func CreateMessage(c *gin.Context) {
	var req dtos.CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channelIDStr := c.Param("channelId")
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userIDStr := userID.(string)

	message, err := services.CreateMessage(req.Content, channelIDStr, userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dtos.ToMessageResponse(*message)

	websockets.ChatHub.Broadcast <- response

	c.JSON(http.StatusCreated, response)
}


// GetMessagesByChannelID godoc
// @Summary        Get all messages in a channel
// @Description    Retrieve all messages for a specific channel
// @Tags           Message
// @Accept         json
// @Produce        json
// @Param          channelId path string true "Channel ID"
// @Success        200 {array} dtos.MessageResponse "Messages retrieved successfully"
// @Failure        400 {object} map[string]interface{} "Invalid data"
// @Failure        401 {object} map[string]interface{} "Unauthorized"
// @Router /channels/{channelId}/messages [get]
func GetMessagesByChannelID(c *gin.Context) {
	channelIDStr := c.Param("channelId")

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userIDStr := userID.(string)

	messages, err := services.GetMessagesByChannelID(channelIDStr, userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dtos.ToMessageResponses(messages)

	c.JSON(http.StatusOK, response)
}
