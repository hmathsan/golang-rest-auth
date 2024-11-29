package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //Postgres Driver
)

var Db *sql.DB

func InitDB() {
	log.Println("Connecting to database")

	connStr := "postgres://postgres:password@localhost:5432/golang-test?sslmode=disable"
	var err error
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	log.Println("Database connected successfully")
}
