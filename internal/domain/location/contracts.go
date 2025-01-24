package location

import (
	"coc-sync/internal/infrastructure/cocapi"
	"coc-sync/internal/infrastructure/storage"
	"context"
)

type IRepository interface {
	GetLocations(ctx context.Context) ([]Location, error)
	StoreLocations(ctx context.Context) error
}

type Repository struct {
	client *cocapi.Client
	db *storage.DB
}

type IService interface {
	GetLocations(ctx context.Context) ([]Location, error)
}

type Location struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode,omitempty"`
}
