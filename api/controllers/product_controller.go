package controllers

import (
	"VenusX/api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := db.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener productos"})
		return
	}
	c.JSON(200, products)
}
func GetProductsByBarCode(c *gin.Context) {
	barcodeParam := c.Param("barcode")
	barcode, err := strconv.ParseUint(barcodeParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Código de barras inválido"})
		return
	}

	product, err := db.GetProductsBarCode(uint(barcode))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener producto"})
		return
	}
	c.JSON(200, product)
}
