package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Content string    `gorm:"type:text;not null" json:"content"`

	User   User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`

	Channel   Channel   `gorm:"foreignKey:ChannelID" json:"channel,omitempty"`
	ChannelID uuid.UUID `gorm:"type:uuid;not null;index" json:"channelId"`

	CreatedAt time.Time      `gorm:"autoCreateTime;index" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
