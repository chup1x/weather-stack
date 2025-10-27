package handlers

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserRequest struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Age   int    `json:"age"`
	City  string `json:"city"`
}

type RegisterUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetUserRequest struct {
	ID uuid.UUID `param:"id"`
}

type GetUserResponse struct {
	ID         uuid.UUID `json:"id" gorm:"id"`
	Login      string    `json:"login" gorm:"login"`
	Name       string    `json:"name" gorm:"name"`
	Sex        string    `json:"sex" gorm:"sex"`
	Age        int       `json:"age" gorm:"age"`
	City       string    `json:"city" gorm:"city"`
	TelegramID int64     `json:"telegram_id" gorm:"telegram_id"`

	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}
