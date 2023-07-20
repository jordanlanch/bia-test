package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type MeterModel struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Address string    `gorm:"type:varchar(255)"`
}
