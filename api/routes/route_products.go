package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupSearchRoute configura la ruta de búsqueda por código de barras
func SetupSearchRoute(router *gin.Engine) {
	router.GET("/search", controllers.GetProductsByBarCode)
}

// SetupSearchByNameRoute configura la ruta de búsqueda por nombre de producto
func SetupSearchByNameRoute(router *gin.Engine) {
	router.GET("/search/name", controllers.SearchProductsByName)
}
func SetupDeleteProductRoute(router *gin.Engine) {
	router.DELETE("/delete_product/:id", controllers.DeleteProduct)
}

// SetupAddProduct sets up the route for adding a product

func SetupAddProduct(router *gin.Engine) {
	router.POST("/add_product", controllers.CreateProducts)
}

// SetupSearchRoute configura la ruta de búsqueda
func SetupGetAllProductsRoute(router *gin.Engine) {
	router.GET("/getProducts", controllers.GetProducts)
}
