package routes

import (
	"net/http"

	"github.com/go-crud/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
