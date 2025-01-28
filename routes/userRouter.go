package routes

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/handlers"
	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUserHandler).Methods("DELETE")
}
