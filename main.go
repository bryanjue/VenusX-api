package main

import (
	"VenusX/api"
	"VenusX/api/config"
)

func main() {
	// Dbconfiguration
	// connect to the database
	config.ConnectDB()
	db := config.GetDB()
	// start the api rest from package api
	app := api.NewApp()
	// run the app
	app.Run()
}
