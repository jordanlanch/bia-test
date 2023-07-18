package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// table name
const (
	ConsumptionTable = "consumption"
)

type Consumption struct {
	ID                 uuid.UUID
	MeterID            uuid.UUID
	Active             float64
	ReactiveInductive  float64
	ReactiveCapacitive float64
	Exported           float64
	Timestamp          time.Time
}

type ConsumptionRepository interface {
	Save(ctx context.Context, consumption *Consumption) error
	FindById(c context.Context, id uuid.UUID) (*Consumption, error)
	FindMany(c context.Context, pagination *Pagination) ([]*Consumption, error)
	FindBetweenDates(ctx context.Context, start, end time.Time, pagination *Pagination) ([]*Consumption, error)
	FindMonthly(c context.Context, start, end time.Time, meterIDs []uuid.UUID, pagination *Pagination) ([]*Consumption, error)
	FindWeekly(c context.Context, start, end time.Time, meterIDs []uuid.UUID, pagination *Pagination) ([]*Consumption, error)
	FindDaily(c context.Context, start, end time.Time, meterIDs []uuid.UUID, pagination *Pagination) ([]*Consumption, error)
}

type ConsumptionUsecase interface {
	GetConsumptionById(id uuid.UUID) (*Consumption, error)
	GetAllConsumptions(pagination *Pagination) ([]*Consumption, error)
	GetMonthlyConsumptions(start, end time.Time, meterIDs []uuid.UUID, pagination *Pagination) ([]*Consumption, error)
	GetWeeklyConsumptions(start, end time.Time, meterIDs []uuid.UUID, pagination *Pagination) ([]*Consumption, error)
	GetDailyConsumptions(start, end time.Time, meterIDs []uuid.UUID, pagination *Pagination) ([]*Consumption, error)
}
