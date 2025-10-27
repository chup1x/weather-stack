package app

import (
	"context"
	"log"
	"log/slog"

	"github.com/chup1x/weather-stack/internal/config"
	"github.com/chup1x/weather-stack/internal/transport/v1/rest"
)

func MustRunApp() {
	slog.Info("read config...")

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("read config: %s", err.Error())
	}

	log.Println("launching server...")

	server := rest.New()
	if err := server.Start(context.Background(), config); err != nil {
		log.Fatalf("start web server: %s", err.Error())
	}
}
