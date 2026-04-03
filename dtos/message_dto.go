package dtos

import (
	"kordchat-be/models"
	"time"

	"github.com/google/uuid"
)

type CreateMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

type MessageResponse struct {
	ID        uuid.UUID         `json:"id"`
	Content   string            `json:"content"`
	User      UserResponse `json:"user"`
	ChannelID uuid.UUID         `json:"channelId"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func ToMessageResponse(message models.Message) MessageResponse {
	return MessageResponse{
		ID:      message.ID,
		Content: message.Content,
		User: UserResponse{
			ID:       message.User.ID,
			Username: message.User.Username,
			Email:    message.User.Email,
		},
		ChannelID: message.ChannelID,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}

func ToMessageResponses(messages []models.Message) []MessageResponse {
	var responses []MessageResponse
	for _, message := range messages {
		responses = append(responses, ToMessageResponse(message))
	}

	return responses
}
