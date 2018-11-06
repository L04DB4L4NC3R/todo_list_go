package controller

import (
	"context"
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
	if r.Method == http.MethodGet {
		a.temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		f := r.Form
		user, err := model.LoginUser(f.Get("username"), f.Get("password"))
		if err != nil {
			log.Println(err)
		} else {
			ctx := context.WithValue(context.Background(), "username", user.Username)
			r.WithContext(ctx)
			http.Redirect(w, r, "/todo", http.StatusMovedPermanently)
		}
	}
}
