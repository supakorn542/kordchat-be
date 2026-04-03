package controllers

import (
	"net/http"

	"kordchat-be/websockets"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"kordchat-be/dtos"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(c *gin.Context) {
	channelIDStr := c.Param("channelId")
	userID, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userIDStr := userID.(string)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &websockets.Client{
		Hub:       websockets.ChatHub,
		Conn:      conn,
		Send:      make(chan dtos.MessageResponse, 256),
		ChannelID: channelIDStr,
		UserID:    userIDStr,
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

}
