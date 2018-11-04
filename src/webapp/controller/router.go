package controller

import (
	"html/template"
	"net/http"
)

var (
	mytodo todo
	myauth auth
)

func Startup(templates *template.Template) {
	myauth.temp = templates.Lookup("login.html")
	mytodo.temp = templates.Lookup("index.html")
	myauth.router()
	mytodo.router()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
