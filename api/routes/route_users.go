package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupSearchRoute configura la ruta de búsqueda
func SetupGetAllUsersRoute(router *gin.Engine) {
	router.GET("/getUsers", controllers.GetUsers)
}

func SetupDeleteUserRoute(router *gin.Engine) {
	router.DELETE("/delete_user/:id", controllers.DeleteUser)
}

// SetupAddUSer sets up the route for adding a user

func SetupAddUser(router *gin.Engine) {
	router.POST("/add_user", controllers.CreateUsers)
}
