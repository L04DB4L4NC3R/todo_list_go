package controller

import (
	"html/template"
	"log"
	"net/http"
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
		if f.Get("username") == "root" && f.Get("password") == "toor" {
			http.Redirect(w, r, "/todo", http.StatusMovedPermanently)
		} else {
			w.Write([]byte("Wrong username or password"))
		}
	}
}
