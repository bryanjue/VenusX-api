package routes

import (
	"VenusX/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupGetAllSalesRoute(router *gin.Engine) {
	router.GET("/sales", controllers.GetSales)
}

func SetupCreateSaleRoute(router *gin.Engine) {
	router.POST("/sales", controllers.CreateSale)
}

func SetupDeleteSaleRoute(router *gin.Engine) {
	router.DELETE("/delete_sale/:id", controllers.DeleteSale)
}
