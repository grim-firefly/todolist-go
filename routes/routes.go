package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/todolist-go/controllers"
)

var RegisterRoutes = func(router *chi.Mux) {
	router.Get("/", controllers.TodoListIndex)
	router.Get("/todo", controllers.GetAllTodoList)
	router.Post("/todo", controllers.TodoListCreate)
	router.Delete("/todo/{id}", controllers.TodoListDelete)
	router.Get("/todo/{id}", controllers.GetTodoList)
	router.Put("/todo/{id}", controllers.TodoListUpdate)

}
