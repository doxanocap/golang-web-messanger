package main

import (
	"github.com/doxanocap/golang-react/backend/pkg/database"
	"github.com/doxanocap/golang-react/backend/pkg/routes"
)

func main() {
	database.Connect()
	routes.SetupRoutes()
}
