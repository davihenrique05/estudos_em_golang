package main

import (
	"fmt"
	"net/http"
	"storeapp/routes"
)

func main() {

	routes.LoadRoutes()

	port := 8000
	fmt.Println("The server is initialized on port", port)
	http.ListenAndServe(fmt.Sprint(":", port), nil)
}
