package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name         string `json:"name"`        // Name of the client
	Nif          string `json:"nif"`         // legal identification of the client
	Adress       string `json:"adress"`      // adress of the client
	Town         string `json:"town"`        //  town/city of the client
	Province     string `json:"province"`    // province of the client
	Postal_code  uint   `json:"postal_code"` // postal code of the client
	Phone_number uint   `json:"phone_numer"` // phone number of the client
}
