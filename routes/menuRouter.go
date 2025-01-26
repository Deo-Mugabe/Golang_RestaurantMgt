package routes

import (
	"golang-restaurant-mgt/Golang_RestaurantMgt/handlers"

	"github.com/gorilla/mux"
)

func MenuRouter(r *mux.Router) {
	r.HandleFunc("/menus", handlers.GetMenusHandler).Methods("GET")
	r.HandleFunc("/menus/{id:[0-9]+}", handlers.GetMenuHandler).Methods("GET")
	r.HandleFunc("/menus", handlers.CreateMenueHandler).Methods("POST")
	r.HandleFunc("/menus/{id:[0-9]+}", handlers.UpdateMenuHandler).Methods("PUT")
	r.HandleFunc("/menus/{id:[0-9]+}", handlers.DeleteMenuHandler).Methods("DELETE")
}
