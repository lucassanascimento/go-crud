package controllers

import (
	"html/template"
	"net/http"

	"github.com/go-crud/models"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
