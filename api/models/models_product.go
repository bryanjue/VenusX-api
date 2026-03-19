package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name    string  `json:"name"`     // Name of the product
	BarCode uint    `json:"bar_code"` // bar code of the product
	Price   float64 `json:"price"`    // Price of the product
}
