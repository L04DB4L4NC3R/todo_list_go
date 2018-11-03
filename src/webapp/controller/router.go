package controller

import (
	"html/template"
	"net/http"
)

var (
	mytodo todo
)

func Startup(templates *template.Template) {
	mytodo.temp = templates.Lookup("index.html")
	mytodo.router()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
