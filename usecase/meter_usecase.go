package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
)

type MeterUsecase struct {
	meterRepository domain.MeterRepository
	contextTimeout  time.Duration
}

func NewMeterUsecase(meterRepository domain.MeterRepository, timeout time.Duration) domain.MeterUsecase {
	return &MeterUsecase{
		meterRepository: meterRepository,
		contextTimeout:  timeout,
	}
}

func (cu *MeterUsecase) GetMeterById(id uuid.UUID) (*domain.Meter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.meterRepository.FindByID(ctx, id)
}

func (cu *MeterUsecase) GetManyMeters(pagination *domain.Pagination) ([]*domain.Meter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.meterRepository.FindMany(ctx, pagination)
}
