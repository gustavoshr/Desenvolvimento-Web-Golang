package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gopher_store/models"
)

var front = template.Must(template.ParseGlob("Front-end/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", newProducts)
	http.HandleFunc("/insert", insertProducts)
	http.HandleFunc("/delete", deleteProducts)
	http.HandleFunc("/edit", editProducts)
	http.HandleFunc("/update", updateProducts)
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

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price", err)
		}

		convertQuantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity", err)
		}

		models.NewProducts(name, description, convertedPrice, convertQuantityInt)
		http.Redirect(w, r, "/", 301)
	}

}

func updateProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price", err)
		}

		convertQuantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity", err)
		}

		models.UpdateProducts(id, name, description, convertedPrice, convertQuantityInt)
		http.Redirect(w, r, "/", 301)
	}
}

func deleteProducts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Error converting id", err)
		http.Redirect(w, r, "/", 301)
		return
	}

	models.DeleteProducts(id)
	http.Redirect(w, r, "/", 301)
}

func editProducts(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProducts(idProduct)
	front.ExecuteTemplate(w, "Edit", product)
}
