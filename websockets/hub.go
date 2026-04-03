package websockets

import (
	"kordchat-be/dtos"
)

type Hub struct {
	Rooms      map[string]map[*Client]bool
	Broadcast  chan dtos.MessageResponse
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]map[*Client]bool),
		Broadcast:  make(chan dtos.MessageResponse),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Rooms[client.ChannelID]; !ok {
				h.Rooms[client.ChannelID] = make(map[*Client]bool)

			}

			h.Rooms[client.ChannelID][client] = true

		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.ChannelID]; ok {
				delete(h.Rooms[client.ChannelID], client)
				close(client.Send)
			}

		case message := <-h.Broadcast:
			connections := h.Rooms[message.ChannelID.String()]

			for client := range connections {
				client.Send <- message
			}
		}

	}
}


var ChatHub *Hub

func init() {
	ChatHub = NewHub()
	go ChatHub.Run()
}