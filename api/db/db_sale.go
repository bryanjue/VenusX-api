package db

import (
	"VenusX/api/config"
	"VenusX/api/models"
)

func GetAllSales() ([]models.Sale, error) {
	var Sales []models.Sale
	result := config.GetDB().Find(&Sales)
	return Sales, result.Error
}

func AddSales(newSales []models.Sale) error {
	result := config.GetDB().Create(&newSales)
	return result.Error
}

func DeleteSaleByID(id uint) error {
	db := config.GetDB()
	result := db.Unscoped().Delete(&models.Sale{}, id)
	return result.Error
}
