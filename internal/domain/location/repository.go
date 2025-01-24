package location

import (
	"coc-sync/internal/infrastructure/cocapi"
	"context"
	"fmt"
)

type locationResponse struct {
	Items []Location `json:"items"`
}

func NewRepository(client *cocapi.Client) IRepository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) GetLocations(ctx context.Context) ([]Location, error) {
	var response locationResponse
	err := r.client.DoGet(ctx, "/locations", &response)
	if err != nil {
		return nil, fmt.Errorf("fetching locations: %w", err)
	}
	return response.Items, nil
}

func (r *Repository) StoreLocations(ctx context.Context) error {
	return nil
}
