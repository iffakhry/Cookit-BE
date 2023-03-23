package models

import (
	"gorm.io/gorm"
)

type IngredientDetail struct {
	gorm.Model
	IngredientID uint
	Name         string `gorm:"not null;type:VARCHAR(50)"`
	Quantity     int    `gorm:"not null;default:1"`
	Unit         string `gorm:"not null;type:VARCHAR(50)"`
}
