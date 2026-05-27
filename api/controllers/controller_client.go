package controllers

import (
	"VenusX/api/db"
	"VenusX/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetClients(c *gin.Context) {
	clients, err := db.GetAllClients()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener clientes"})
		return
	}
	c.JSON(200, clients)
}

func CreateClients(c *gin.Context) {
	var newClients []models.Client

	if err := c.ShouldBindJSON(&newClients); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	err := db.AddClients(newClients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar los clientes"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Clientes creados exitosamente"})
}

func DeleteClient(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = db.DeleteClientByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al borrar el Cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado correctamente"})
}

/*
	func GetProductsByBarCod(c *gin.Context) {
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

	func SearchProductByBarCoe(c *gin.Context) {
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
*/
