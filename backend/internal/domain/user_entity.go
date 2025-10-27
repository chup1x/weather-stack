package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type UserID struct {
	uuid.UUID
}

type UserEntity struct {
	ID         UserID `json:"id" gorm:"id"`
	Login      string `json:"login" gorm:"login"`
	Name       string `json:"name" gorm:"name"`
	Sex        string `json:"sex" gorm:"sex"`
	Age        int    `json:"age" gorm:"age"`
	City       string `json:"city" gorm:"city"`
	TelegramID int64  `json:"telegram_id" gorm:"telegram_id"`

	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}
