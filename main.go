package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/routes"
	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Initialize the database connection
	db.InitDB()
	// Create a new router
	r := mux.NewRouter()

	// Register routes for each model
	routes.FoodRouter(r)
	routes.InvoiceRouter(r)
	routes.MenuRouter(r)
	routes.OrderItemRouter(r)
	routes.OrderRouter(r)
	routes.TableRouter(r)
	routes.UserRouter(r)

	// Start the server
	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
