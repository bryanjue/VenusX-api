package api

import (
	"VenusX/api/routes"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	// Crea una nueva instancia de Gin
	router := gin.Default()

	// Configura las rutas
	routes.SetupHelloRoute(router)
	routes.SetupSearchRoute(router)

	return &App{
		router: router,
	}
}

func (a *App) Run() {
	// Ejecuta la aplicación en el puerto 8080
	a.router.Run(":8080")
}
