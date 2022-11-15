package database

import "time"

type Comment struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	UserId    int            `json:"user_id" gorm:"not null"`
	PhotoId   int            `json:"photo_id" gorm:"not null"`
	Message   string         `json:"message" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	User      *UserResponse  `json:"User,omitempty"`
	Photo     *PhotoResponse `json:"Photo,omitempty"`
}
