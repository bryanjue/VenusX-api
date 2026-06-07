package db

import (
	"VenusX/api/config"
	"VenusX/api/models"
)

func GetAllSales() ([]models.Sale, error) {
	var sales []models.Sale
	result := config.GetDB().Order("created_at desc").Find(&sales)
	return sales, result.Error
}

func CreateSale(sale *models.Sale) error {
	result := config.GetDB().Create(sale)
	return result.Error
}

func DeleteSaleByID(id uint) error {
	db := config.GetDB()
	result := db.Unscoped().Delete(&models.Sale{}, id)
	return result.Error
}
