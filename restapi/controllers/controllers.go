package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/database"
	"restapi/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllCelebrities(w http.ResponseWriter, r *http.Request) {
	var p []models.Celebrity

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func GetCelebrityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var celebrity models.Celebrity
	database.DB.First(&celebrity, id)
	json.NewEncoder(w).Encode(celebrity)
}
