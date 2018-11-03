package templ

import (
	"html/template"
)

func PopulateTemplates() *template.Template {
	t := template.New("index.html")
	template.Must(t.ParseGlob("static/*.html"))
	return t
}
