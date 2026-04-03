package websockets

import (
	"github.com/gorilla/websocket"
	"kordchat-be/dtos"
)

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan dtos.MessageResponse
	ChannelID string
	UserID    string
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Send
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		err := c.Conn.WriteJSON(message)
		if err != nil {
			return
		}
	}


}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_,_,err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
