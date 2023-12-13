package main

import (
	"fmt"
	"restapi/database"
	"restapi/routes"
)

func main() {

	database.ConnectWithDatabase()

	fmt.Println("Initializing server...")
	routes.HandleRequest()
}
