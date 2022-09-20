package main

import (
	"net/http"

	"github.com/go-crud/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
