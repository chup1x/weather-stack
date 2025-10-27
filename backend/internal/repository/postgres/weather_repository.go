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

func (r *WetherRepository) GetWeatherRequestByUserID(ctx context.Context, id domain.UserID) ([]*domain.WeatherRequestEntity, error) {
	weatherRequests := []*domain.WeatherRequestEntity{}

	if err := r.db.WithContext(ctx).Table("weather_requests").Find(&weatherRequests).Error; err != nil {
		return nil, err
	}

	return weatherRequests, nil
}
