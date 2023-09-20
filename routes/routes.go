package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/todolist-go/controllers"
)

var RegisterRoutes = func(router *chi.Mux) {
	router.Get("/", controllers.TodoListIndex)
}
