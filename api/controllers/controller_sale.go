package controllers

import (
	"VenusX/api/db"
	"VenusX/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSales(c *gin.Context) {
	sales, err := db.GetAllSales()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener productos"})
		return
	}
	c.JSON(200, sales)
}

/*func GetProductsByBarCode(c *gin.Context) {
	barcodeParam := c.Query("Bar_code")
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
}*/

func CreateSales(c *gin.Context) {
	var newSales []models.Sale

	if err := c.ShouldBindJSON(&newSales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	err := db.AddSales(newSales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar la venta"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Productos creados exitosamente"})
}

func DeleteSale(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = db.DeleteSaleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al borrar la venta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Venta eliminado correctamente"})
}
