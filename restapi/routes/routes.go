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
	r.HandleFunc("/api/celebrities", controllers.GetAllCelebrities).Methods("Get")
	r.HandleFunc("/api/celebrities/{id}", controllers.GetCelebrityById).Methods("Get")
	r.HandleFunc("/api/celebrities", controllers.CreatNewCelebrity).Methods("Post")
	r.HandleFunc("/api/celebrities/{id}", controllers.DeleteCelebrity).Methods("Delete")
	r.HandleFunc("/api/celebrities/{id}", controllers.UpdateCelebrity).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
