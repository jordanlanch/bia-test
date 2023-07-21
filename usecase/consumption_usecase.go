package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
)

type consumptionUsecase struct {
	consumptionRepository domain.ConsumptionRepository
	contextTimeout        time.Duration
}

func NewConsumptionUsecase(consumptionRepository domain.ConsumptionRepository, timeout time.Duration) domain.ConsumptionUsecase {
	return &consumptionUsecase{
		consumptionRepository: consumptionRepository,
		contextTimeout:        timeout,
	}
}

func (cu *consumptionUsecase) GetConsumptionByID(id uuid.UUID) (*domain.Consumption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindByID(ctx, id)
}

func (cu *consumptionUsecase) GetConsumptionsByPeriod(periodType, start, end string, meterID []int, pagination *domain.Pagination) (*domain.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindByPeriod(ctx, periodType, start, end, meterID, pagination)
}
