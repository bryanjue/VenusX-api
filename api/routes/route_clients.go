package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

/*
	SetupSearchRoute configura la ruta de búsqueda

	func SetupSearchRoute(router *gin.Engine) {
		router.GET("/search", controllers.GetProductsByBarCode)
	}

// SetupAddClient sets up the route for adding a product
*/
func SetupAddClient(router *gin.Engine) {
	router.POST("/add_client", controllers.CreateClients)
}

// SetupSearchRoute configura la ruta de búsqueda
func SetupGetAllClientRoute(router *gin.Engine) {
	router.GET("/getclients", controllers.GetClients)
}
func SetupDeleteClientRoute(router *gin.Engine) {
	router.DELETE("/delete_client/:id", controllers.DeleteClient)
}
