package usercntrl

import (
	"github.com/chup1x/weather-stack/internal/domain"
	"github.com/google/uuid"
)

type GetTelegramProfileRequest struct {
	ID int64 `param:"telegram_id" validate:"required"`
}

type GetTelegramProfileResponse struct {
	*domain.UserEntity
}

type GetProfileRequest struct {
	ID domain.UserID `param:"user_id" validate:"required"`
}

type GetProfileResponse struct {
	*domain.UserEntity
}

type RegisterProfileRequest struct {
	*domain.UserEntity
}

type RegisterProfileResponse struct {
	ID uuid.UUID `json:"id"`
}
