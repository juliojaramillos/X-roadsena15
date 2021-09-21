package database

import (
	"database/sql"
	"log"

	"unal.edu.co/rest-example/config"
)

var DB *sql.DB

func InitDatabase() {
	switch config.Database.Dialect {
	case "postgres":
		initPostgres()
	default:
		log.Fatalf("Selected dialect is not implemented\n")
	}
}
