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
	var celebrities []models.Celebrity

	database.DB.Find(&celebrities)

	json.NewEncoder(w).Encode(celebrities)
}

func GetCelebrityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var celebrity models.Celebrity
	database.DB.First(&celebrity, id)
	json.NewEncoder(w).Encode(celebrity)
}

func CreatNewCelebrity(w http.ResponseWriter, r *http.Request) {
	var newCelebrity models.Celebrity
	json.NewDecoder(r.Body).Decode(&newCelebrity)
	database.DB.Create(&newCelebrity)
	json.NewEncoder(w).Encode(newCelebrity)
}

func DeleteCelebrity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var celebrity models.Celebrity
	database.DB.Delete(&celebrity, id)
	json.NewEncoder(w).Encode(celebrity)
}

func UpdateCelebrity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var celebrity models.Celebrity
	database.DB.First(&celebrity, id)
	json.NewDecoder(r.Body).Decode(&celebrity)
	database.DB.Save(&celebrity)
	json.NewEncoder(w).Encode(celebrity)
}
