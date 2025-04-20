package routes

import (
	"fmt"
	"net/http"
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
}
