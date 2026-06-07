package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	Items      string          `json:"items" gorm:"type:jsonb"`      // JSON of cart items
	Subtotal   float64         `json:"subtotal"`
	TaxRate    float64         `json:"tax_rate"`
	TaxAmount  float64         `json:"tax_amount"`
	Discount   float64         `json:"discount"`
	Total      float64         `json:"total"`
	Method     string          `json:"method"`       // Payment method (cash, card, etc.)
	Cashier    string          `json:"cashier"`
	Status     string          `json:"status"`       // completed, refunded
	Comment    string          `json:"comment"`
	RawItems   json.RawMessage `json:"-" gorm:"-"`   // Helper for JSON parsing
}
