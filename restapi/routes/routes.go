package routes

import (
	"log"
	"net/http"
	"restapi/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/celebrities", controllers.GetAllCelebrities).Methods("Get")
	r.HandleFunc("/celebrities/{id}", controllers.GetCelebrityById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
