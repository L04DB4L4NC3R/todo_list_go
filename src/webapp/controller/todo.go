package controller

import (
	"html/template"
	"net/http"

	"../model"
)

type todo struct {
	temp *template.Template
}

func (t todo) router() {
	http.HandleFunc("/", t.handleTodo)
}

func (t todo) handleTodo(w http.ResponseWriter, r *http.Request) {
	vm := model.Newdata()
	t.temp.Execute(w, vm)
}
