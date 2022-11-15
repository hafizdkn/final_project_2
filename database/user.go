package database

import (
	"time"
)

type User struct {
	ID        int              `json:"id" gorm:"primaryKey"`
	Username  string           `json:"username" gorm:"not null"`
	Email     string           `json:"email" gorm:"unique;not null"`
	Password  string           `json:"password" gorm:"not null"`
	Age       int              `json:"age" gorm:"not null"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Photos    *[]PhotoResponse `json:"Photo,omitempty" gorm:"constraint:OnDelete:SET NULL"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (UserResponse) TableName() string {
	return "users"
}
