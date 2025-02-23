package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/services"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/utility"
)

// GetUsersHandler retrieves a list of all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

// GetUserHandler retrieves a specific user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	user, err := services.GetUser(uint(id))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "User not found")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// CreateUserHandler handles user registration
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	err := services.CreateUser(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// UpdateUserHandler handles updating user details
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	err = services.UpdateUser(uint(id), &user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	respondWithJSON(w, http.StatusAccepted, user)
}

// DeleteUserHandler deletes a user by ID
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	err = services.DeleteUser(uint(id))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

// SignUp handles user registration
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Check if the email or phone already exists
	if err := services.CreateUser(&user); err != nil {
		respondWithError(w, http.StatusConflict, err.Error())
		return
	}

	// Generate JWT tokens
	accessToken, refreshToken, err := services.GenerateTokens(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating tokens")
		return
	}

	user.Token = accessToken
	user.RefreshToken = refreshToken

	// Respond with the user data and tokens
	respondWithJSON(w, http.StatusCreated, user)
}

// Login handles user authentication and token generation
func Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid inputs")
		return
	}

	// Get user by email
	user, err := services.GetUserByEmail(loginData.Email)
	if err != nil || user == nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Verify password
	valid, err := services.VerifyPassword(user.Password, loginData.Password)
	if !valid {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate new tokens
	token, refreshToken, err := services.GenerateTokens(*user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating tokens")
		return
	}

	// Respond with tokens
	respondWithJSON(w, http.StatusOK, map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

// Utility function to respond with JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Utility function to respond with error message
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
