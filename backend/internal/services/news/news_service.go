package services

import (
	"context"
	"fmt"

	"github.com/chup1x/weather-stack/internal/domain"
)

type newsCreater interface {
	CreateNewsRequest(context.Context, *domain.NewsEntity) error
}

type newsProvider interface {
	GetNewsByCityID(context.Context, string) (*domain.NewsEntity, error)
}

type newsStorage interface {
	newsCreater
	newsProvider
}

type NewsService struct {
	repo     newsStorage
	newsServ *newsClient
}

func NewNewsService(repo newsStorage, news *newsClient) *NewsService {
	return &NewsService{
		repo:     repo,
		newsServ: news,
	}
}

func (s *NewsService) GetNews(ctx context.Context, cityID string) (map[string]any, error) {
	news, err := s.newsServ.GetNewsByCityID(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("to get a news by city id - %s: %w", cityID, err)
	}

	return news, nil
}
