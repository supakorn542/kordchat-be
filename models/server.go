package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Server struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name    string    `gorm:"type:varchar(100);not null" json:"name"`
	OwnerID uuid.UUID `gorm:"type:uuid;not null;index" json:"ownerId"`

	Channels []Channel `gorm:"foreignKey:ServerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"channels,omitempty"`
	Users    []User    `gorm:"many2many:server_users;" json:"users,omitempty"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
