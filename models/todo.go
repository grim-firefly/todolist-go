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

// deleting todo
func DeleteTodoList(id interface{}) {
	db := database.GetDB()
	db.Delete(&Todo{}, id)

}

// get single row
func GetTodo(id interface{}) Todo {
	db := database.GetDB()
	var todo Todo
	db.First(&todo, id)
	return todo
}

// update todo list

func UpdateTodo(id interface{}, data *Todo) Todo {
	db := database.GetDB()
	var todo Todo
	db.First(&todo, id)
	todo.Title = data.Title
	todo.IsActive = data.IsActive
	db.Save(&todo)
	return todo
}
