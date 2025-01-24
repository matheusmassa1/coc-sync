package location

import (
	"context"
	"fmt"
)

type service struct {
	repo IRepository
}

func NewService(repo IRepository) IService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetLocations(ctx context.Context) ([]Location, error) {
	locations, err := s.repo.GetLocations(ctx)
	if err != nil {
		return nil, fmt.Errorf("service failed to fetch locations: %w", err)
	}

	return locations, nil
}
