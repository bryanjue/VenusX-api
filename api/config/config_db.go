package config

import (
	"VenusX/api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"user_tpv",
		"12345678",
		"tpv",
		"5432",
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&models.Product{})
	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.Order{})
	DB.Debug().AutoMigrate(&models.Product{})

	fmt.Println("Database connection successfully established")
}

func GetDB() *gorm.DB {

	return DB
}
