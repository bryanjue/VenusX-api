package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// Ruta GET
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "¡Hola, mundo!"})
	})

	// Iniciar servidor en el puerto 8080
	app.Run(":8080")
}
