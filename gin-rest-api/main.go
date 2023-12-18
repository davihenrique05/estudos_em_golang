package main

import (
	"gin-rest-api/database"
	"gin-rest-api/routes"
)

// @title Studente API
// @version 1.0
func main() {
	database.ConnectWithDatabase()

	routes.HandleRequests()
}
