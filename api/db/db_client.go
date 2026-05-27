package db

import (
	"VenusX/api/config"
	"VenusX/api/models"
)

func GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	result := config.GetDB().Find(&clients)
	return clients, result.Error
}

func AddClients(newClients []models.Client) error {
	result := config.GetDB().Create(&newClients)
	return result.Error
}

func DeleteClientByID(id uint) error {
	db := config.GetDB()
	result := db.Unscoped().Delete(&models.Client{}, id)
	return result.Error
}

/*
	func GetProductsBarCode(barcode uint) (*models.Product, error) {
		var product models.Product
		result := config.GetDB().Where("bar_code = ?", barcode).First(&product)
		if result.Error != nil {
			return nil, result.Error
		}
		return &product, nil
	}
*/
