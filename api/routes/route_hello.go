package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupHelloRoute(router *gin.Engine) {
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
}
