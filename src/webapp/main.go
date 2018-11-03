package main

import (
	"log"
	"net/http"

	"./controller"
	"./view"
)

func main() {
	templates := templ.PopulateTemplates()
	controller.Startup(templates)
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
