package app

import (
	"coc-sync/internal/config"
	"coc-sync/internal/domain/location"
	"coc-sync/internal/infrastructure/cocapi"
	"net/http"
)

type App struct {
	Location location.IService
	client   *http.Client
}

func NewApp(cfg config.Config) (*App, error) {
	httpClient := &http.Client{
		Timeout: cfg.HTTPTimeout,
	}

	cocClient := cocapi.NewClient(httpClient, cfg.APIKey)

	// Repos
	locationRepo := location.NewRepository(cocClient)

	return &App{
		Location: location.NewService(locationRepo),
		client:   httpClient,
	}, nil
}

func (a *App) Close() error {
	a.client.CloseIdleConnections()
	return nil
}
