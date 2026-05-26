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

/*func GetProductsBarCode(barcode uint) (*models.Product, error) {
	var product models.Product
	result := config.GetDB().Where("bar_code = ?", barcode).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func AddProducts(newProducts []models.Product) error {
	result := config.GetDB().Create(&newProducts)
	return result.Error
}

func DeleteProductByID(id uint) error {
	db := config.GetDB()
	result := db.Unscoped().Delete(&models.Product{}, id)
	return result.Error
}
*/
