package weathercntrl

import (
	"github.com/chup1x/weather-stack/internal/repository/postgres"
	newsservice "github.com/chup1x/weather-stack/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterNewsRoutes(router fiber.Router, db *gorm.DB) {
	newsRepo := postgres.NewNewsRepository(db)
	newsCntrl := NewNewsController(newsservice.NewNewsService(newsRepo))

	news := router.Group("/news")
	news.Get("/:city_id", newsCntrl.GetNewsHandler)
}
