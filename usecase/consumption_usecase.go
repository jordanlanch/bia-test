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

func (cu *consumptionUsecase) GetConsumptionById(id uuid.UUID) (*domain.Consumption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindById(ctx, id)
}

func (cu *consumptionUsecase) GetConsumptionsByPeriod(period_type, start, end string, meterID []int, pagination *domain.Pagination) (*domain.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindByPeriod(ctx, period_type, start, end, meterID, pagination)
}
