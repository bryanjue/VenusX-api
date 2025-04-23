package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAddProduct sets up the route for adding a product

func SetupAddProduct(router *gin.Engine) {
	router.POST("/add_product", controllers.CreateProducts)
}
