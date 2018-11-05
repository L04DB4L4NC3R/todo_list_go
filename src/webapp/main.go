package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"./controller"
	"./middleware"
	"./model"
	"./view"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	templates := templ.PopulateTemplates()
	db := connectToDB()
	defer db.Close()
	controller.Startup(templates)
	log.Fatalln(http.ListenAndServe(":3000", &middleware.TimeoutMiddleware{
		Next: new(middleware.GzipCompression)}))
}

func connectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		log.Fatalln(fmt.Errorf("Error connecting to DB"))
	}
	model.StartupDB(db)
	return db
}
