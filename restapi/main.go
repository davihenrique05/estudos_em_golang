package main

import (
	"fmt"
	"restapi/models"
	"restapi/routes"
)

func main() {
	models.Celebrities = []models.Celebrity{
		{Name: "Nome 1", Biography: "Very sad background"},
		{Name: "Nome 2", Biography: "Very sad background"},
	}

	fmt.Println("Initializing server...")
	routes.HandleRequest()
}
