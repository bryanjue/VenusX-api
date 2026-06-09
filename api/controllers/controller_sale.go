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
		c.JSON(500, gin.H{"error": "Error al obtener ventas"})
		return
	}
	c.JSON(200, sales)
}

func CreateSale(c *gin.Context) {
	var newSales []models.Sale

	if err := c.ShouldBindJSON(&newSales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	err := db.AddSale(newSales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar la venta"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Venta creada exitosamente"})
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

	c.JSON(http.StatusOK, gin.H{"message": "Venta eliminada correctamente"})
}
