package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/platinumscatter/port-service/internal/domain"
)

type PortService struct {
}

func NewPortService() PortService {
	return PortService{}
}

func (s PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	randomID := uuid.New().String()
	return domain.NewPort(randomID, randomID, randomID, randomID, randomID,
		[]string{randomID}, []string{randomID}, []float64{1.0, 2.0}, randomID, randomID, nil)
}
