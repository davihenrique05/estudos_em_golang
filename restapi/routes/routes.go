package routes

import (
	"log"
	"net/http"
	"restapi/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/celebrities", controllers.GetAllCelebrities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
