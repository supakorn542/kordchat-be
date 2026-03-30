package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255); not null" json:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Messages []Message `gorm:"foreignKey:UserID" json:"messages,omitempty"`
	Servers  []Server  `gorm:"many2many:server_users;" json:"servers,omitempty"`
}
