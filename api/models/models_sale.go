package models

import (
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ClientID *uint  `json:"client_id"`
	Client   Client `json:"client" gorm:"foreignKey:ClientID"`

	Subtotal  float64 `json:"subtotal"`
	TaxRate   float64 `json:"taxRate"`
	TaxAmount float64 `json:"taxAmount"`
	Discount  float64 `json:"discount"`

	Total         float64 `json:"total"`
	PaymentMethod string  `json:"payment_method"`
	Cashier       string  `json:"cashier"`
	Status        string  `json:"status"`
	Comment       string  `json:"comment"`

	Items []SaleItem `json:"items" gorm:"foreignKey:SaleID"`
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
