package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver (blank identifier to avoid error
)

var (
	db    *sql.DB
	dbErr error
)

func main() {
	// docker run -it -d -p 5432:5432 --name cities-postgre -e POSTGRES_PASSWORD=cities-pass -d postgres:alpine3.14

	// Database connection parameters
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "cities-pass"
	dbname := "godb"

	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, dbErr = sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		panic(dbErr)
	}

	insertCity()
}

/*
	Create -> Insert
	Read   -> Select
	Update -> Update
	Delete -> Delete
*/

// Insert city to database table cities (name, code)
func insertCity() {
	r, err := db.Exec("insert into cities (name, code) values ('istanbul', 34)") // Exec returns sql.Result
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(r.RowsAffected()) // Returns number of rows affected by the query
	}
}
