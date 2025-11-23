package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/chup1x/weather-stack/internal/domain"
)

type weatherCreater interface {
	CreateWeatherRequest(context.Context, *domain.WeatherEntity) error
}

type weatherProvider interface {
	GetWeatherByCity(context.Context, string) (*domain.WeatherEntity, error)
	//GetClothesByComb(context.Context, int) (*domain.WeatherClothesEntity, error)
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

func (s *WeatherService) CreateWeatherRecord(ctx context.Context, new *domain.WeatherEntity) error {
	if err := s.repo.CreateWeatherRequest(ctx, new); err != nil {
		return fmt.Errorf("to create a weatcher request: %w", err)
	}
	return nil
}

func (s *WeatherService) GetWeather(ctx context.Context, city string) (*domain.WeatherEntity, error) {
	//news_en := &domain.NewsEntity{}

	baseURL := "https://newsapi.org/v2/everything"
	params := url.Values{}
	params.Add("q", "Санкт-Петербург")
	params.Add("from", "2025-11-09")
	params.Add("sortBy", "publishedAt")
	params.Add("language", "ru")
	params.Add("apiKey", "0fac40f7dcd34967af176019e1c6a526")

	fullURL := baseURL + "?" + params.Encode()

	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// _, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := r.db.WithContext(ctx).Table("news").Where("city_id = ?", city).First(news_en).Error; err != nil {

	// 	// filename := fmt.Sprintf("temp_news_%s.json", city)
	// 	// err = os.WriteFile(".json", body, 0644)
	// 	// if err != nil {
	// 	// 	log.Fatal("Error writing file:", err)
	// 	// }
	// 	news, _ := json.Marshal(body)
	// 	// news_en := &domain.NewsEntity{
	// 	// 	PATH: filename,
	// 	// }

	// 	// if err := r.db.WithContext(ctx).Table("news").Create(news_en).Error; err != nil {
	// 	// 	log.Fatal("Error writing file to database:", err)
	// 	// 	return
	// 	// }
	// 	return news, nil
	// }

	weather, err := s.repo.GetWeatherByCity(ctx, city)
	if err != nil {
		return nil, fmt.Errorf("to select a weather by city: %w", err)
	}
	return weather, nil
}

/*
	func (s *WeatherService) GetWeatherClothes(ctx context.Context, id int64) (*domain.WeatherClothesEntity, error) {
		//comb, err := s.repo.SelectByTelegramID(ctx, id)
		clothes, err := s.repo.GetClothesByComb(ctx, 1234)
		if err != nil {
			return nil, fmt.Errorf("to select a weather by city: %w", err)
		}
		return clothes, nil
	}
*/
