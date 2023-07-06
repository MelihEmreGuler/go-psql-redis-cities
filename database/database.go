package database

import (
	"database/sql"
	"fmt"
)

var (
	DB    *sql.DB
	dbErr error
)

func Connect() {

	// Database connection parameters
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "cities-pass"
	dbname := "godb"

	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) // Sprintf returns string
	DB, dbErr = sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		panic(dbErr)
	}
}
