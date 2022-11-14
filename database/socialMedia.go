package database

import (
	"time"
)

type SocialMedia struct {
	ID             int           `json:"id" gorm:"primaryKey"`
	Name           string        `json:"name" gorm:"not null"`
	SocialMeidaUrl string        `json:"social_media_url" gorm:"not null"`
	UserId         int           `json:"user_id" gorm:"not null"`
	User           *UserResponse `json:"User,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}
