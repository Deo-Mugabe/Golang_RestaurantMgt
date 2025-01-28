package routes

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/handlers"
	"github.com/gorilla/mux"
)

func OrderItemRouter(r *mux.Router) {
	r.HandleFunc("/orderItems", handlers.GetOrderItemsHandler).Methods("GET")
	r.HandleFunc("/orderItems/{id:[0-9]+}", handlers.GetOrderItemHandler).Methods("GET")
	r.HandleFunc("/orderItems", handlers.CreateOrderItemHandler).Methods("POST")
	r.HandleFunc("/orderItems/{id:[0-9]+}", handlers.UpdateOrderItemHandler).Methods("PUT")
	r.HandleFunc("/orderItems/{id:[0-9]+}", handlers.DeleteOrderItemHandler).Methods("DELETE")
	r.HandleFunc("/orderItems-order/{order_id:[0-9]}", handlers.GetOrderItemsByOrderHandler).Methods("GET")
}
