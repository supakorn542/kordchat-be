package dtos

import(
	"github.com/google/uuid"
	"time"

	"kordchat-be/models"
)


type ServerResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uuid.UUID `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateServerRequest struct {
	Name string `json:"name" binding:"required"`
}


func ToServerResponse(server models.Server) ServerResponse {
	return  ServerResponse{
		ID:        server.ID,
		Name:      server.Name,
		OwnerID:   server.OwnerID,
		CreatedAt: server.CreatedAt,
		UpdatedAt: server.UpdatedAt,
	}
}


func ToServerResponses(servers []models.Server) []ServerResponse {
	var responses []ServerResponse
	for _, server := range servers {
		responses = append(responses, ToServerResponse(server))
	}

	return  responses
}