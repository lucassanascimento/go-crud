package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("template/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Produto{
		{"Camisa", "Preta", 100.0, 10},
		{"Camisa", "Preta", 100.0, 10},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
