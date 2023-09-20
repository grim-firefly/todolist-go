package config

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/todolist-go/routes"
)

// route registering function
func RegisterRoutes() {
	chi := chi.NewRouter()
	routes.RegisterRoutes(chi)
	err := http.ListenAndServe(":8080", chi)
	if err != nil {
		panic(err)
	}
}
