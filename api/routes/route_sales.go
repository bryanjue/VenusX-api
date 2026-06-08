package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAddSale(router *gin.Engine) {
	router.POST("/add_sale", controllers.CreateSales)
}

// SetupSearchRoute configura la ruta de búsqueda
func SetupGetAllSaleRoute(router *gin.Engine) {
	router.GET("/getSales", controllers.GetSales)
}
func SetupDeleteSaleRoute(router *gin.Engine) {
	router.DELETE("/delete_Sale/:id", controllers.DeleteSale)
}
