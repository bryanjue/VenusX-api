package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupSearchRoute configura la ruta de búsqueda
func SetupSearchRoute(router *gin.Engine) {
	router.GET("/search", controllers.GetProductsByBarCode)
}
