package dtos

type CreateServerRequest struct {
	Name string `json:"name" binding:"required"`
}