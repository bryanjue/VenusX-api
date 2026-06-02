package db

import (
	"VenusX/api/config"
	"VenusX/api/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.GetDB().Find(&users)
	return users, result.Error
}

func AddUsers(newUsers []models.User) error {
	result := config.GetDB().Create(&newUsers)
	return result.Error
}

func DeleteUserByID(id uint) error {
	db := config.GetDB()
	result := db.Unscoped().Delete(&models.Product{}, id)
	return result.Error
}

/*func GetProductsBarCode(barcode uint) (*models.Product, error) {
	var product models.Product
	result := config.GetDB().Where("bar_code = ?", barcode).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

*/
