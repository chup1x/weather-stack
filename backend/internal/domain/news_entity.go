package domain

import (
	"errors"
	"time"
)

var NewsNotFound = errors.New("news not found")

type NewsEntity struct {
	CityID    string    `json:"city_id" gorm:"city_id"`
	Path      string    `json:"path" gorm:"path"`
	CreatedAt time.Time `json:"-" gorm:"created_at"`
}
