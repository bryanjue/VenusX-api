package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

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
