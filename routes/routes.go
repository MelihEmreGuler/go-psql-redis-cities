package routes

import (
	"github.com/MelihEmreGuler/go-psql-redis-cities/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHandlers(r *mux.Router) {

	r.HandleFunc("/city", handlers.GetCity).Methods(http.MethodGet)
	r.HandleFunc("/city", handlers.PostCity).Methods(http.MethodPost)
	r.HandleFunc("/city", handlers.PutCity).Methods(http.MethodPut)
	r.HandleFunc("/city", handlers.DeleteCity).Methods(http.MethodDelete)
}
