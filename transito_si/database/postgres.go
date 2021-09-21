package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"

	"unal.edu.co/rest-example/config"
)

func initPostgres() {
	psqlStringConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Database, getSSLMode(),
	)

	var err error
	log.Printf("Connecting to database...\n%s\n", psqlStringConn)
	DB, err = sql.Open(config.Database.Dialect, psqlStringConn)
	if err != nil {
		log.Fatalf("Connection to %s database [%s] failed:\n%v,\n", config.Database.Dialect, config.Database.Database, err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging to %s database [%s]:\n%v\n", config.Database.Dialect, config.Database.Database, err)
	}

	log.Printf("Connected to the %s database [%s]\n", config.Database.Dialect, config.Database.Database)

}

func getSSLMode() string {
	if config.Database.SSLMode {
		return "enable"
	} else {
		return "disable"
	}
}
