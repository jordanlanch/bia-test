package domain

import (
	"context"

	"github.com/google/uuid"
)

// table name
const (
	ConsumptionTable = "consumptions"
)

type Consumption struct {
	ID                 uuid.UUID `json:"id"`
	MeterID            int       `json:"meter_id"`
	Active             float64   `json:"active"`
	ReactiveInductive  float64   `json:"reactive_inductive"`
	ReactiveCapacitive float64   `json:"reactive_capacitive"`
	Exported           float64   `json:"exported"`
	Period             string    `json:"period"`
}

type Response struct {
	Period    []string                 `json:"period"`
	DataGraph []map[string]interface{} `json:"data_graph"`
}

type ConsumptionRepository interface {
	Save(ctx context.Context, consumption *Consumption) error
	FindById(c context.Context, id uuid.UUID) (*Consumption, error)
	FindByPeriod(c context.Context, period_type, start, end string, meterIDs []int, pagination *Pagination) (*Response, error)
}

type ConsumptionUsecase interface {
	GetConsumptionById(id uuid.UUID) (*Consumption, error)
	GetConsumptionsByPeriod(period_type, start, end string, meterIDs []int, pagination *Pagination) (*Response, error)
}
