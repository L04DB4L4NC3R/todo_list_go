package main

import (
	"net/http"

	"./view"
)

func main() {
	templates := templ.PopulateTemplates()
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Test"))
		res := templates.Lookup("index.html")
		if res != nil {
			templates.Execute(w, nil)

		}
	})
	http.ListenAndServe(":3000", nil)
}
