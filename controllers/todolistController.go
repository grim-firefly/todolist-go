package controllers

import (
	"encoding/json"
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

func TodoListCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		title := r.Form.Get("title")
		is_active := r.PostForm.Get("title")
		ResponseJson(w, 200, struct {
			Title    string `json:"title"`
			IsActive string `json:"is_active"`
		}{
			Title:    title,
			IsActive: is_active,
		})
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
