package routes

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/handlers"
	"github.com/gorilla/mux"
)

func FoodRouter(r *mux.Router) {
	r.HandleFunc("/foods", handlers.GetFoodsHandler).Methods("GET")
	r.HandleFunc("/foods/{id:[0-9]+}", handlers.GetFoodHandler).Methods("GET")
	r.HandleFunc("/foods", handlers.CreateFoodHandler).Methods("POST")
	r.HandleFunc("/foods/{id:[0-9]+}", handlers.UpdateFoodHandler).Methods("PUT")
	r.HandleFunc("/foods/{id:[0-9]+}", handlers.DeleteFoodHandler).Methods("DELETE")
}
