package main

import (
	"VenusX/api"
)

func main() {
	// start the api rest from package api
	app := api.NewApp()
	// run the app
	app.Run()
}
