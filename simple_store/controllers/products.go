package controllers

import (
	"html/template"
	"log"
	"net/http"
	"simple_store/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("ERROR  while converting price: ", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("ERROR while converting quantitiy: ", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DelectProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
