package postgres

import (
	"context"
	"errors"
	"net/http"
    "net/url"
	"log"
	"io"
	"fmt"
	"os"

	"github.com/chup1x/weather-stack/internal/domain"
	"gorm.io/gorm"
)

type WeatherRepository struct {
	db *gorm.DB
}

func NewWeatherRepository(db *gorm.DB) *WeatherRepository {
	return &WeatherRepository{db: db}
}

func (r *WeatherRepository) CreateWeatherRequest(ctx context.Context, new *domain.WeatherEntity) error {
	return r.db.WithContext(ctx).Table("weather_requests").Create(new).Error
}

func (r *WeatherRepository) GetWeatherByCity(ctx context.Context, city string) (*domain.WeatherEntity, error) {
	weather := &domain.WeatherEntity{}

	if err := r.db.WithContext(ctx).Table("weather_requests").Where("city_id = ?", city).First(weather).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return weather, nil
}
/*
func (r *WeatherRepository) GetClothesByComb(ctx context.Context, id int) ([]*domain.WeatherClothesEntity, error) {
	clothes := []*domain.WeatherClothesEntity{}

	if err := r.db.WithContext(ctx).Table("clothes").Where("id = ?", id).First(clothes).Error; err != nil {
		return nil, err
	}

	return clothes, nil
}
*/
func (r *WeatherRepository) GetNewsByCity(ctx context.Context, city string) (*domain.NewsEntity, error) {
	news := &domain.NewsEntity{}

	if err := r.db.WithContext(ctx).Table("news").Where("city_id = ?", city).First(news).Error; err != nil {
	

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

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("temp_news_%s.json", city)
		err = os.WriteFile(".json", body, 0644)
		if err != nil {
			log.Fatal("Error writing file:", err)
		}
		
		news := &domain.NewsEntity{
			PATH: filename,
		}

		// if err := r.db.WithContext(ctx).Table("news").Create(news).Error; err != nil {
		// 	log.Fatal("Error writing file to database:", err)
		return news, nil
		// }
	}

	return news, nil
}
