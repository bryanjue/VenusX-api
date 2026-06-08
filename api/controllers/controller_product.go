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
	barcodeParam := c.Query("Bar_code")
	barcode, err := strconv.ParseUint(barcodeParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Código de barras inválido"})
		return
	}

	product, err := db.GetProductsBarCode(uint(barcode))
	if err != nil {
		c.JSON(404, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(200, product)
}

func SearchProductsByName(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de búsqueda no puede estar vacío"})
		return
	}

	products, err := db.GetProductsByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar productos"})
		return
	}

	c.JSON(http.StatusOK, products)
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

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = db.DeleteProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al borrar el producto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
