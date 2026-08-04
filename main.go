package main

import (
	"html/template"
	"net/http"

	"gopher_store/models"
)

var front = template.Must(template.ParseGlob("Front-end/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", newProducts)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	front.ExecuteTemplate(w, "Index", allProducts)
}

func newProducts(w http.ResponseWriter, r *http.Request) {
	front.ExecuteTemplate(w, "new-products", nil)
}

func insertProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")
	}

}
