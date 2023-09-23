package config

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/todolist-go/routes"
)

// route registering function
func RegisterRoutes() {
	chi := chi.NewRouter()
	// setup ui path like css, js
	SetupUiPath(chi)

	// registing all logical routes
	routes.RegisterRoutes(chi)

	// serving
	err := http.ListenAndServe(":8080", chi)
	if err != nil {
		panic(err)
	}

}

// setup ui path
func SetupUiPath(chi *chi.Mux) {
	fs := http.FileServer(http.Dir("public"))
	chi.Get("/public/*", http.StripPrefix("/public/", fs).ServeHTTP)
}
