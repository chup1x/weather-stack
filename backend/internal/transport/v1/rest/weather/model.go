package weathercntrl

type CreateWeatherRecord struct {
	TelegramID  int64   `json:"telegram_id"`
	Temperature float64 `json:"temperature"`
	Clothing    string  `json:"clothing"`
}
