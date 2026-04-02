package dtos

import (
	"github.com/google/uuid"
	"time"

	"kordchat-be/models"
)

type CreateChannelRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type"`
}

type ChannelResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	ServerID  uuid.UUID `json:"serverId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToChannelResponse(channel models.Channel) ChannelResponse {
	return ChannelResponse{
		ID:        channel.ID,
		Name:      channel.Name,
		Type:      channel.Type,
		ServerID:  channel.ServerID,
		CreatedAt: channel.CreatedAt,
		UpdatedAt: channel.UpdatedAt,
	}
}

func ToChannelResponses(channels []models.Channel) []ChannelResponse {
	var responses []ChannelResponse
	for _, channel := range channels {
		responses = append(responses, ToChannelResponse(channel))
	}
	return responses
}
