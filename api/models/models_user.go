package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"` // Name of the client
	Password string `json:"-"`    // legal identification of the client
	Role     uint   `json:"role"` // postal code of the client
}
