package services

import (
	"context"
	"fmt"

	"github.com/chup1x/weather-stack/internal/domain"
)

type newsProvider interface {
	GetNewsByCityID(context.Context, string) (*domain.NewsEntity, error)
}

type newsStorage interface {
	newsProvider
}

type NewsService struct {
	repo newsStorage
}

func NewNewsService(repo newsStorage) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) GetNews(ctx context.Context, cityID string) (*domain.NewsEntity, error) {
	news, err := s.repo.GetNewsByCityID(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("to select a weather by city: %w", err)
	}
	return news, nil
}
