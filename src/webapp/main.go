package main

import (
	"log"
	"net/http"

	"./controller"
	"./middleware"
	"./view"
)

func main() {
	templates := templ.PopulateTemplates()
	controller.Startup(templates)
	log.Fatalln(http.ListenAndServe(":3000", &middleware.TimeoutMiddleware{
		Next: new(middleware.GzipCompression)}))
}
