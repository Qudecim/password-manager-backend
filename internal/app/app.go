package app

import "database/sql"

var Config *ApplicationConfig
var DB *sql.DB

type ApplicationConfig struct {
	Debug bool
}

func New(database *sql.DB) {
	config := ApplicationConfig{true}
	Config = &config

	DB = database
}
