package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Channel struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Server   Server    `gorm:"foreignKey:ServerID" json:"server,omitempty"`
	ServerID uuid.UUID `gorm:"type:uuid;not null;index" json:"serverId"`

	Messages []Message `gorm:"foreignKey:ChannelID" json:"messages,omitempty"`
}
