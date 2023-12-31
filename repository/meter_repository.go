package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
	"gorm.io/gorm"
)

type meterRepository struct {
	db    *gorm.DB
	table string
}

func NewMeterRepository(db *gorm.DB, table string) domain.MeterRepository {
	return &meterRepository{db, table}
}

func (r *meterRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Meter, error) {
	var meter domain.Meter

	result := r.db.WithContext(ctx).First(&meter, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &meter, nil
}

func (r *meterRepository) FindMany(ctx context.Context, pagination *domain.Pagination) ([]*domain.Meter, error) {
	var meters []*domain.Meter
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	result := db.Find(&meters)
	if result.Error != nil {
		return nil, result.Error
	}
	return meters, nil
}
