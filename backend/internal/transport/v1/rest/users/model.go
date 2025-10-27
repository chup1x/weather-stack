package usercntrl

import (
	"github.com/chup1x/weather-stack/internal/domain"
)

type GetTelegramProfileRequest struct {
	ID int64 `param:"telegram_id" validate:"required"`
}

type GetTelegramProfileResponse struct {
	*domain.UserEntity
}

type GetProfileRequest struct {
	ID domain.UserID `json:"user_id" validate:"required"`
}

type GetProfileResponse struct {
	*domain.UserEntity
}

type RegisterProfileRequest struct {
	Login      string `json:"login,omitempty"`
	Name       string `json:"name" validate:"required"`
	Sex        string `json:"sex" validate:"required"`
	Age        int    `json:"age" validate:"required"`
	City       string `json:"city" validate:"required"`
	TelegramID int64  `json:"telegram_id,omitempty"`
}

type RegisterProfileResponse struct {
	ID domain.UserID `json:"id"`
}
