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
	// Ruta GET /search?Bar_code=...
	router.GET("/search", func(c *gin.Context) {
		query := c.Query("Bar_code") // Obtiene el parámetro "Bar_code" de la URL
		code, name := strconv.Atoi(query)
		var results []struct {
			ID       int
			Bar_code int
			Name     string
		}
		found := false
		// Busca coincidencias en la lista
		for _, article := range articles {
			if article.Bar_code == code {
				found = true
				break // Termina el bucle si lo encuentra
			}
		}
		message := fmt.Sprintf("No se encontraron resultados para %d %d elementos", name, code)
		// Devuelve los resultados en JSON
		if found {
			c.JSON(http.StatusOK, gin.H{
				"results": results,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": message,
			})
		}
	})
}
