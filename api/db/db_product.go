package db

import (
	"VenusX/api/config"
	"VenusX/api/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := config.GetDB().Find(&products)
	return products, result.Error
}
func GetProductsBarCode(barcode uint) (*models.Product, error) {
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
