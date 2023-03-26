package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"main:main@tcp(127.0.0.1:3306)/main")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	rows, err := db.Query("select PersonID, LastName from Persons where PersonID = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
