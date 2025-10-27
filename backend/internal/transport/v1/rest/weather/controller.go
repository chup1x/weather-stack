package weathercntrl

import (
	"github.com/chup1x/weather-stack/internal/domain"
	weatherservice "github.com/chup1x/weather-stack/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type weatherController struct {
	s         *weatherservice.WeatherService
	validator *validator.Validate
}

func NewWeatherController(service *weatherservice.WeatherService) *weatherController {
	return &weatherController{
		validator: validator.New(),
		s:         service,
	}
}

func (cn *weatherController) CreateWeatherRecordHandler(c *fiber.Ctx) error {
	var req CreateWeatherRecord
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	if err := cn.validator.Struct(req); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if err := cn.s.CreateWeatherRecord(c.UserContext(), &domain.WeatherRequestEntity{
		Temperature: req.Temperature,
		Clothing:    req.Clothing,
	}); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (cn *weatherController) GetWeatherHistoryHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
