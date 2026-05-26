package config

import (
	"VenusX/api/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file loaded, using existing environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&models.Product{})
	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.Order{})
	DB.Debug().AutoMigrate(&models.Product{})
	/*DB.AutoMigrate(&models.Client{})
	DB.Debug().AutoMigrate(&models.Client{})*;/

	fmt.Println("Database connection successfully established")
}

func GetDB() *gorm.DB {

	return DB
}
