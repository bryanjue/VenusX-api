package routes

import (
	"VenusX/api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupSearchRoute configura la ruta de búsqueda
func SetupSearchRoute(router *gin.Engine) {
	router.GET("/search", func(c *gin.Context) {
		query := c.Query("Bar_code")

		// Verificar si el parámetro está vacío
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "El código de barras no puede estar vacío",
			})
			return
		}

		// Intentar convertir el código de barras a entero
		code, err := strconv.Atoi(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("El código de barras '%s' no es válido", query),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Código de barras recibido: %d", code),
		})
	})

	router.POST("/products", func(c *gin.Context) {
		var newProducts []models.Product

		if err := c.ShouldBindJSON(&newProducts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
			return
		}

		// Leer el archivo news_products.json
		file, err := os.OpenFile("api/news_products.json", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al abrir el archivo"})
			return
		}
		defer file.Close()

		var products []models.Product

		// Leer los productos existentes
		if err := json.NewDecoder(file).Decode(&products); err != nil {
			products = []models.Product{}
		}

		// Agregar los nuevos productos
		products = append(products, newProducts...)

		// Guardar los productos en el archivo
		file.Truncate(0)
		file.Seek(0, 0)
		if err := json.NewEncoder(file).Encode(products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar el producto"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Producto creado exitosamente"})
	})
}
