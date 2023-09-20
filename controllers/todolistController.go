package controllers

import (
	"html/template"
	"net/http"
)

func TodoListIndex(writer http.ResponseWriter, request *http.Request) {
	parseTemplate, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}
	parseTemplate.Execute(writer, nil)

}
