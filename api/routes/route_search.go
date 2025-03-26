package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Lista de ejemplo (simulando una base de datos)
var articles = []struct {
	ID       int
	Bar_code int
	Name     string
}{
	{1, 1234567890123, "CocaCola"},
}

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

		var results []struct {
			ID       int
			Bar_code int
			Name     string
		}
		found := false

		// Buscar coincidencias en la lista
		for _, article := range articles {
			if article.Bar_code == code {
				found = true
				results = append(results, article)
				break
			}
		}

		// Responder con los resultados o mensaje de error
		if found {
			c.JSON(http.StatusOK, gin.H{
				"results": results,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("No se encontraron resultados para %d", code),
			})
		}
	})
}
