package routes

import (
	"golang-restaurant-mgt/Golang_RestaurantMgt/handlers"

	"github.com/gorilla/mux"
)

func TableRouter(r *mux.Router) {
	r.HandleFunc("/tables", handlers.GetBooksHandler).Methods("GET")
	r.HandleFunc("/tables/{id:[0-9]+}", handlers.GetBookHandler).Methods("GET")
	r.HandleFunc("/tables", handlers.CreateBookHandler).Methods("POST")
	r.HandleFunc("/tables/{id:[0-9]+}", handlers.UpdateBookHandler).Methods("PUT")
	r.HandleFunc("/tables/{id:[0-9]+}", handlers.DeleteBookHandler).Methods("DELETE")
}
