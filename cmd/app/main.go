package main

import (
	"database/sql"
	"log"

	"github.com/qudecim/password-manager-backend/internal/transport/rest"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qudecim/password-manager-backend/internal/app"
)

// Let's go mthfck
func main() {

	db, err := sql.Open("mysql",
		"main:main@tcp(127.0.0.1:3306)/main")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app.New(db)
	rest.Run()

}
