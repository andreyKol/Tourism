package app

import (
	"log/slog"
	"net/http"
	"os"

	"Tourism/internal/common/config"
)

type App struct {
	httpServer *http.Server
}

func New(configPath string) (*App, error) {
	config.MustReadConfigFromFile(configPath)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	return &App{}, nil
}
