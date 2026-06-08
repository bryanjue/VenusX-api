package models

import (
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ClientID      *uint      `json:"client_id"`
	Client        Client     `json:"client" gorm:"foreignKey:ClientID"`
	Total         float64    `json:"total"`
	PaymentMethod string     `json:"payment_method"`
	Status        string     `json:"status"`
	Items         []SaleItem `json:"items" gorm:"foreignKey:SaleID"`
}

type SaleItem struct {
	gorm.Model
	SaleID    uint    `json:"sale_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Subtotal  float64 `json:"subtotal"`
}
