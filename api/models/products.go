package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`     // Name of the product
	Bar_code uint    `json:"bar_code"` // bar code of the product
	Price    float64 `json:"price"`    // Price of the product
}
