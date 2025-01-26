package routes

import (
	"golang-restaurant-mgt/Golang_RestaurantMgt/handlers"

	"github.com/gorilla/mux"
)

func OrderRouter(r *mux.Router) {
	r.HandleFunc("/orders", handlers.GetOrdersHandler).Methods("GET")
	r.HandleFunc("/orders/{id:[0-9]+}", handlers.GetOrderHandler).Methods("GET")
	r.HandleFunc("/orders", handlers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{id:[0-9]+}", handlers.UpdateOrderHandler).Methods("PUT")
	r.HandleFunc("/orders/{id:[0-9]+}", handlers.DeleteOrderHandler).Methods("DELETE")
}
