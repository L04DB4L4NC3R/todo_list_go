package model

import "database/sql"

var db *sql.DB

func StartupDB(database *sql.DB) {
	db = database
}
