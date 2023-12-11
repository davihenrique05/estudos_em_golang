package main

import (
	"fmt"
	"net/http"
	"simple_store/routes"
)

func main() {

	routes.LoadRoutes()

	porta := 8000
	fmt.Println("Servidor inicilizado em porta", porta)
	http.ListenAndServe(fmt.Sprint(":", porta), nil)
}
