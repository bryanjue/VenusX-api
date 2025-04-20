package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupSearchRoute configura la ruta de búsqueda
func SetupGetAllProductsRoute(router *gin.Engine) {
	router.GET("/getproducts", controllers.GetProducts)
}
