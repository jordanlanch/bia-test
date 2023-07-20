package domain

import (
	"context"

	"github.com/google/uuid"
)

const (
	MeterTable = "meters"
)

type Meter struct {
	ID      uuid.UUID
	Address string
}

type MeterRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*Meter, error)
	FindMany(ctx context.Context, pagination *Pagination) ([]*Meter, error)
}

type MeterUsecase interface {
	GetMeterById(id uuid.UUID) (*Meter, error)
	GetManyMeters(pagination *Pagination) ([]*Meter, error)
}
