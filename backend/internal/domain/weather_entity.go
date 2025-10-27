package domain

import (
	"time"
)

type WeatherRequestEntity struct {
	ID          int64     `json:"id" gorm:"id"`
	UserID      UserID    `json:"user_id" gorm:"user_id"`
	Temperature float64   `json:"temperature" gorm:"temperature"`
	Clothing    string    `json:"clothing" gorm:"clothing"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
