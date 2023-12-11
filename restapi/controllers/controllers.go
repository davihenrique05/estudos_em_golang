package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllCelebrities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Celebrities)
}
