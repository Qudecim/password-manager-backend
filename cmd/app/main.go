package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qudecim/password-manager-backend/internal/app"
	service "github.com/qudecim/password-manager-backend/internal/services"
)

func init() {

	db, err := sql.Open("mysql",
		"main:main@tcp(127.0.0.1:3306)/main")
	if err != nil {
		log.Fatal(err)
	}

	service.DB = db
}

// Let's go mthfck
func main() {

	app.Run()

}
