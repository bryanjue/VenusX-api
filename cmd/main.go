package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Estructura para representar un producto
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Slice para almacenar los productos en memoria
var products []Product
var nextID = 1

func main() {
	// Crea una instancia del router
	router := gin.Default()

	// Ruta de ejemplo: GET /
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "¡Hola, mundo!",
		})
	})

	// Ruta para datos de ejemplo
	router.GET("/api/datos", func(c *gin.Context) {
		data := map[string]interface{}{
			"nombre": "Mi API",
			"estado": "funcionando",
		}
		c.JSON(http.StatusOK, data)
	})

	// Ruta para listar todos los productos
	router.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	// Ruta para agregar un nuevo producto
	router.POST("/api/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err == nil {
			newProduct.ID = nextID
			nextID++
			products = append(products, newProduct)
			c.JSON(http.StatusCreated, newProduct)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Ruta para actualizar un producto existente
	router.PUT("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id) // Convertir id a int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err == nil {
			for i, product := range products {
				if product.ID == intID { // Comparar con el ID convertido
					products[i] = updatedProduct
					products[i].ID = product.ID // Mantener el ID original
					c.JSON(http.StatusOK, products[i])
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Ruta para eliminar un producto
	router.DELETE("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id) // Convertir id a int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		for i, product := range products {
			if product.ID == intID { // Comparar con el ID convertido
				products = append(products[:i], products[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
	})

	// Inicia el servidor en el puerto 8081
	router.Run(":8081")
}
