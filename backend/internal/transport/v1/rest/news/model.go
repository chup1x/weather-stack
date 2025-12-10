package weathercntrl

import (
	"github.com/chup1x/weather-stack/internal/domain"
)

type GetNewsRequest struct {
	City string `params:"city" validate:"required"`
}

type GetNewsByTelegramRequest struct {
	TelegramID int `params:"telegram_id" validate:"required"`
}

type GetNewsResponse struct {
	*domain.NewsEntity
}
