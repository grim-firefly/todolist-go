package models

import (
	"github.com/grim-firefly/todolist-go/database"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title    string `json:"title"`
	IsActive bool   `json:"is_active"`
}

// creating todolist
func CreateTodo(todo *Todo) *Todo {
	db := database.GetDB()
	db.Create(&todo)
	return todo
}

// getting all todolist
func GetAllTodoList() []Todo {
	db := database.GetDB()
	var todolists []Todo
	db.Find(&todolists)
	return todolists
}
