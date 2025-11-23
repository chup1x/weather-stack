package weathercntrl

import (
	"github.com/chup1x/weather-stack/internal/domain"
)

type GetWeatherHistoryRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetWeatherHistoryResponse struct {
	*domain.WeatherEntity
}

type GetWeatherClothesRequest struct {
	User int64 `json:"telegram_id" validate:"required"`
}

type GetWeatherClothesResponse struct {
	*domain.WeatherClothesEntity
}

type CreateWeatherRequest struct {
	*domain.WeatherEntity
}

type CreateWeatherResponse struct {
	City string `json:"city"`
}
