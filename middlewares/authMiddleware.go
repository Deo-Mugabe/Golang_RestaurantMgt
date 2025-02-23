package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/helpers"
)

// AuthMiddleware is a function that validates the JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		// Extract the token from the Authorization header
		// Expected format: Bearer <token>
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			respondWithError(w, http.StatusUnauthorized, "Invalid Authorization format")
			return
		}

		// Validate the token
		user, err := helpers.ValidateToken(tokenString[1])
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid token: %s", err.Error()))
			return
		}

		// Attach the user to the request context (optional, if needed in the handlers)
		// You can use context to pass the user to your handlers if needed
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)
		r = r.WithContext(ctx)

		// Proceed with the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}

// Utility function to respond with error message in JSON format
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
