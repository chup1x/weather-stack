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
	ID         uuid.UUID `json:"id" gorm:"id"`
	Name       string    `json:"name" gorm:"name"`
	Sex        string    `json:"sex" gorm:"sex"`
	Age        int       `json:"age" gorm:"age"`
	CityN      string    `json:"city_n" gorm:"city_n"`
	CityW      string    `json:"city_w" gorm:"city_w"`
	DropTime   string    `json:"drop_time" gorm:"drop_time"`
	TComfort   int       `json:"t_comfort" gorm:"t_comfort"`
	TTol       int       `json:"t_tol" gorm:"t_tol"`
	TPuh       int       `json:"t_puh" gorm:"t_puh"`
	Temp1      int       `json:"temp1" gorm:"temp1"`
	Temp2      string    `json:"-" gorm:"temp2"`
	Password   string    `json:"password" gorm:"password"`
	TelegramID int64     `json:"telegram_id" gorm:"telegram_id"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
}
