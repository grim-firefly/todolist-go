package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

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
