package routes

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/handlers"
	"github.com/gorilla/mux"
)

func TableRouter(r *mux.Router) {
	r.HandleFunc("/tables", handlers.GetTablesHandler).Methods("GET")
	r.HandleFunc("/tables/{id:[0-9]+}", handlers.GetTableHandler).Methods("GET")
	r.HandleFunc("/tables", handlers.CreateTableHandler).Methods("POST")
	r.HandleFunc("/tables/{id:[0-9]+}", handlers.UpdateTableHandler).Methods("PUT")
	r.HandleFunc("/tables/{id:[0-9]+}", handlers.DeleteTableHandler).Methods("DELETE")
}
