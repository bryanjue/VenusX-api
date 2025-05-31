package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupSearchRoute configura la ruta de búsqueda
func SetupSearchRoute(router *gin.Engine) {
	router.GET("/search", controllers.GetProductsByBarCode)
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
	router.GET("/getproducts", controllers.GetProducts)
}
