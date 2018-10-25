package main

import (
	"html/template"
	"log"
	"net/http"
)

var items = []string{"milk", "eggs", "ham"}

func main() {
	myMux := http.NewServeMux()
	templ := genTemp()
	myMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path[1:]
		t := templ.Lookup(file + ".html")
		if t != nil {
			err := t.Execute(w, items)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	log.Fatal(http.ListenAndServe(":3000", myMux))
}

func genTemp() *template.Template {
	t := template.New("index.html")
	template.Must(t.ParseGlob("views/*.html"))
	return t
}
