package weathercntrl

type GetWeatherHistoryRequest struct {
	ID int `param:"city_w" validate:"required"`
}

type GetWeatherHistoryResponse struct {
	*domain.WeatherEntity
}

type GetWeatherClothesRequest struct {
	ID string `json:"combination" validate:"required"`
}

type GetWeatherClothesResponse struct {
	*domain.WeatherClothesEntity
}

type GetNewsRequest struct {
	ID string `json:"city_n" validate:"required"`
}

type GetNewsResponse struct {
	*domain.NewsEntity
}

type CreateWeatherRecord struct {
	TelegramID  int64   `json:"telegram_id"`
	Temperature float64 `json:"temperature"`
	Clothing    string  `json:"clothing"`
}
