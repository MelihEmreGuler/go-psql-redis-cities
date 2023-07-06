package main

import (
	"fmt"
	"github.com/MelihEmreGuler/go-psql-redis-cities/database"
	"github.com/MelihEmreGuler/go-psql-redis-cities/repository"
	"github.com/MelihEmreGuler/go-psql-redis-cities/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // Postgres driver (blank identifier to avoid error)
	"net/http"
)

func main() {
	//Connect to database
	database.Connect()

	// Create a new repository (singleton pattern) (only one instance of this struct)
	repository.NewRepo(database.DB)

	// Create a new router
	r := mux.NewRouter()

	// Register handlers
	routes.RegisterHandlers(r)

	// goroutine to run http server
	go func() {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// main goroutine will wait here
	<-make(chan struct{})

}
