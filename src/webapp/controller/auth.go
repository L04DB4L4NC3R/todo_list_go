package controller

import (
	"html/template"
	"log"
	"net/http"

	"../model"
)

type auth struct {
	temp *template.Template
}

func (a auth) router() {
	http.HandleFunc("/", a.HandleAuth)
}

func (a auth) HandleAuth(w http.ResponseWriter, r *http.Request) {
	user := model.NewUserType()
	if r.Method == http.MethodGet {
		a.temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		f := r.Form
		user, err = model.LoginUser(f.Get("username"), f.Get("password"))
	}
}
