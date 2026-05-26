package controllers

import (
	"VenusX/api/db"

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

/*func GetProductsByBarCod(c *gin.Context) {
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

func CreateProdcts(c *gin.Context) {
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

func DeleteProdct(c *gin.Context) {
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
*/
