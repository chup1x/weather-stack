package services

import (
	"context"
	"fmt"

	"github.com/chup1x/weather-stack/internal/domain"
)

type weatherCreater interface {
	CreateWeatherRequest(context.Context, *domain.WeatherRequestEntity) error
}

type weatherProvider interface {
	GetWeatherRequestByUserID(context.Context, domain.UserID) ([]*domain.WeatherRequestEntity, error)
}

type weatherStorage interface {
	weatherCreater
	weatherProvider
}

type WeatherService struct {
	repo weatherStorage
}

func NewWeatherService(repo weatherStorage) *WeatherService {
	return &WeatherService{
		repo: repo,
	}
}

func (s *WeatherService) CreateWeatherRecord(ctx context.Context, new *domain.WeatherRequestEntity) error {
	if err := s.repo.CreateWeatherRequest(ctx, new); err != nil {
		return fmt.Errorf("to create a weatcher request: %w", err)
	}
	return nil
}
