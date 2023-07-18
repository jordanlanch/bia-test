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

func (cu *consumptionUsecase) GetAllConsumptions(pagination *domain.Pagination) ([]*domain.Consumption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindMany(ctx, pagination)
}

func (cu *consumptionUsecase) GetMonthlyConsumptions(start, end time.Time, meterID []uuid.UUID, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindMonthly(ctx, start, end, meterID, pagination)
}

func (cu *consumptionUsecase) GetWeeklyConsumptions(start, end time.Time, meterID []uuid.UUID, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindWeekly(ctx, start, end, meterID, pagination)
}

func (cu *consumptionUsecase) GetDailyConsumptions(start, end time.Time, meterID []uuid.UUID, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.consumptionRepository.FindDaily(ctx, start, end, meterID, pagination)
}
