package controllers

import (
	"VenusX/api/db"
	"VenusX/api/models"
	"net/http"
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

func SearchProductByBarCode(c *gin.Context) {
	query := c.Query("Bar_code")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El código de barras no puede estar vacío",
		})
		return
	}

	code, err := strconv.Atoi(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El código de barras no es válido",
		})
		return
	}

	product, err := db.GetProductsBarCode(uint(code))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener el producto",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProducts(c *gin.Context) {
	var newProducts []models.Product

	if err := c.ShouldBindJSON(&newProducts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	err := db.AddProducts(newProducts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar los productos"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Productos creados exitosamente"})
}
