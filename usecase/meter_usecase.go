package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
)

type meterUsecase struct {
	meterRepository domain.MeterRepository
	contextTimeout  time.Duration
}

func NewMeterUsecase(meterRepository domain.MeterRepository, timeout time.Duration) domain.MeterUsecase {
	return &meterUsecase{
		meterRepository: meterRepository,
		contextTimeout:  timeout,
	}
}

func (cu *meterUsecase) GetMeterById(id uuid.UUID) (*domain.Meter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.meterRepository.FindById(ctx, id)
}

func (cu *meterUsecase) GetManyMeters(pagination *domain.Pagination) ([]*domain.Meter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.meterRepository.FindMany(ctx, pagination)
}
