package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/todolist-go/models"
)

// index page
func TodoListIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parseTemplate, err := template.ParseFiles("view/index.html")
		if err != nil {
			panic(err)
		}
		parseTemplate.Execute(w, nil)
	} else if r.Method == "POST" {
		res := models.GetAllTodoList()
		ResponseJson(w, 200, res)
	}

}

// get all todo list
func GetAllTodoList(w http.ResponseWriter, r *http.Request) {
	res := models.GetAllTodoList()
	ResponseJson(w, 200, res)
}

// create todolist
func TodoListCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		title := r.PostForm.Get("title")
		is_active := r.PostForm.Get("is_active")
		// data := make(map[string]interface{})
		todo := models.Todo{
			Title:    title,
			IsActive: is_active == "true",
		}
		models.CreateTodo(&todo)

		ResponseJson(w, 200, todo)
		log.Println("Written in TodoList Create")
	}
	// return "hi"
}

// delete
func TodoListDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := chi.URLParam(r, "id")
		models.DeleteTodoList(id)
		ResponseJson(w, 200, struct {
			message string
		}{
			message: "Deleted SuccessFully",
		})
	}
}

// info of single todo
func GetTodoList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := chi.URLParam(r, "id")
		todo := models.GetTodo(id)
		ResponseJson(w, 200, todo)
	}
}

// update todolist
// create todolist
func TodoListUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		r.ParseForm()
		id := chi.URLParam(r, "id")
		title := r.PostForm.Get("title")
		is_active := r.PostForm.Get("is_active")
		// data := make(map[string]interface{})
		todo := models.Todo{
			Title:    title,
			IsActive: is_active == "true",
		}
		updatedTodo := models.UpdateTodo(id, &todo)

		ResponseJson(w, 200, updatedTodo)
		log.Println("Written in TodoList Update")
	}
	// return "hi"
}

// response writter
func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

// responser json
func ResponseJson(w http.ResponseWriter, code int, data interface{}) {
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}
