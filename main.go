package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/middlewares"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
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

	// Protect routes with AuthMiddleware
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middlewares.AuthMiddleware)
	protectedRoutes.HandleFunc("/protected", ProtectedHandler).Methods("GET")

	// Start the server
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Access the authenticated user from the context
	user := r.Context().Value("user").(*models.User)

	// Respond with the authenticated user's information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "This is a protected route",
		"user":    user.Email,
	})
}
