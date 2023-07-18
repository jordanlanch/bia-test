package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
	"gorm.io/gorm"
)

type consumptionRepository struct {
	db    *gorm.DB
	table string
}

func NewConsumptionRepository(db *gorm.DB, table string) domain.ConsumptionRepository {
	return &consumptionRepository{db, table}
}

func (r *consumptionRepository) Save(ctx context.Context, consumption *domain.Consumption) error {
	result := r.db.WithContext(ctx).Save(consumption)
	return result.Error
}

func (r *consumptionRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Consumption, error) {
	var consumption domain.Consumption

	result := r.db.WithContext(ctx).First(&consumption, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &consumption, nil
}

func (r *consumptionRepository) FindMany(ctx context.Context, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	var consumptions []*domain.Consumption
	db := r.db.WithContext(ctx)

	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}

	result := r.db.Find(&consumptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return consumptions, nil
}

func (r *consumptionRepository) FindBetweenDates(ctx context.Context, start, end time.Time, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	var consumptions []*domain.Consumption
	db := r.db.WithContext(ctx)

	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}
	result := r.db.WithContext(ctx).Where("created_at BETWEEN ? AND ?", start, end).Find(&consumptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return consumptions, nil
}

func (r *consumptionRepository) FindMonthly(ctx context.Context, start, end time.Time, meterIDs []uuid.UUID, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	var consumptions []*domain.Consumption
	db := r.db.WithContext(ctx)

	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}

	result := db.WithContext(ctx).Where("meter_id IN (?) AND created_at BETWEEN ? AND ?", meterIDs, start, end).Find(&consumptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return consumptions, nil
}

func (r *consumptionRepository) FindWeekly(ctx context.Context, start, end time.Time, meterIDs []uuid.UUID, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	var consumptions []*domain.Consumption
	db := r.db.WithContext(ctx)

	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}

	result := db.WithContext(ctx).Where("meter_id IN (?) AND created_at BETWEEN ? AND ?", meterIDs, start, end).Find(&consumptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return consumptions, nil
}

func (r *consumptionRepository) FindDaily(ctx context.Context, start, end time.Time, meterIDs []uuid.UUID, pagination *domain.Pagination) ([]*domain.Consumption, error) {
	var consumptions []*domain.Consumption

	db := r.db.WithContext(ctx)

	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}

	result := db.WithContext(ctx).Where("meter_id IN (?) AND created_at BETWEEN ? AND ?", meterIDs, start, end).Find(&consumptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return consumptions, nil
}
