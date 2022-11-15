package database

import "time"

type Photo struct {
	ID        int           `json:"id" gorm:"primaryKey"`
	Title     string        `json:"title" gorm:"not null"`
	Caption   string        `json:"caption" `
	PhotoUrl  string        `json:"photo_url" gorm:"not null"`
	UserId    int           `json:"user_id" gorm:"not null"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserResponse `json:"User,omitempty"`
}

type PhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption" `
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

func (PhotoResponse) TableName() string {
	return "photos"
}
