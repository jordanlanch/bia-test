package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Email    string    `gorm:"type:varchar(255)"`
	Password string    `gorm:"type:varchar(255)"`
}
