package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ConsumptionModel struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;primary_key"`
	MeterID            uuid.UUID `gorm:"type:uuid"`
	Active             float64   `gorm:"type:int"`
	ReactiveInductive  float64   `gorm:"type:float"`
	ReactiveCapacitive float64   `gorm:"type:float"`
	Exported           float64   `gorm:"type:float"`
	Timestamp          time.Time `gorm:"type:timestamp"`
}
