package controller

import (
	"html/template"
	"log"
	"net/http"

	"../model"
)

type todo struct {
	temp *template.Template
}

func (t todo) router() {
	http.HandleFunc("/todo", t.handleTodo)
}

func (t todo) handleTodo(w http.ResponseWriter, r *http.Request) {
	vm := model.Newdata()
	log.Println(vm.Todo)
	t.temp.Execute(w, vm)
}
