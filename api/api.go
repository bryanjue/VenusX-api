package api

import (
	"VenusX/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	// Crea una nueva instancia de Gin
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Configura las rutas
	//routes of products
	routes.SetupSearchRoute(router)
	routes.SetupGetAllProductsRoute(router)
	routes.SetupAddProduct(router)
	routes.SetupDeleteProductRoute(router)

	//routes of Clients
	routes.SetupGetAllClientRoute(router)
	routes.SetupAddClient(router)
	routes.SetupDeleteClientRoute(router)

	//routes of Users
	routes.SetupGetAllUsersRoute(router)
	routes.SetupAddUser(router)
	routes.SetupDeleteUserRoute(router)

	return &App{
		router: router,
	}
}

func (a *App) Run() {
	// Ejecuta la aplicación en el puerto 8080
	a.router.Run(":8080")
}
