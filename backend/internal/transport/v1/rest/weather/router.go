package weathercntrl

import (
	"github.com/chup1x/weather-stack/internal/repository/postgres"
	weatherservice "github.com/chup1x/weather-stack/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterWeatherRoutes(router fiber.Router, db *gorm.DB) {
	weatherRepo := postgres.NewWeatherRepository(db)
	weatherCntrl := NewWeatherController(weatherservice.NewWeatherService(weatherRepo))

	weather := router.Group("/weather")

	weather.Get("/:id", weatherCntrl.GetWeatherHandler)
	weather.Post("/", weatherCntrl.CreateWeatherRecordHandler)

	//weather.Get("/clothes", weatherCntrl.GetWeatherClothesHandler)
}
