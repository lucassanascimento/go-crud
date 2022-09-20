package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-crud/models"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)
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

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro de conversão de preço: ", err)
		}
		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro de conversão de quantidade: ", err)
		}

		models.CreateProduct(name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}
