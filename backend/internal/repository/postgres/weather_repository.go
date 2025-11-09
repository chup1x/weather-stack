package postgres

import (
	"context"

	"github.com/chup1x/weather-stack/internal/domain"
	"gorm.io/gorm"
)

type WetherRepository struct {
	db *gorm.DB
}

func NewWetherRepository(db *gorm.DB) *WetherRepository {
	return &WetherRepository{db: db}
}

func (r *WetherRepository) CreateWeatherRequest(ctx context.Context, new *domain.WeatherRequestEntity) error {
	return r.db.WithContext(ctx).Table("weather_requests").Create(new).Error
}

func (r *WetherRepository) GetWeatherByCity(ctx context.Context, city string) ([]*domain.WeatherEntity, error) {
	weather := []*domain.WeatherEntity{}

	if err := r.db.WithContext(ctx).Table("weather_requests").Find(&weather).Error; err != nil {
		return nil, err
	}

	return weather, nil
}
