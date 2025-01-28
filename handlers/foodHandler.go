package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/services"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/utility"
)

// Get all foods
func GetFoodsHandler(w http.ResponseWriter, r *http.Request) {
	foods, err := services.GetFoods()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, foods)
}

// Get single food
func GetFoodHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	food, err := services.GetFood(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	utility.JsonResponse(w, http.StatusOK, food)
}

// Create a food
func CreateFoodHandler(w http.ResponseWriter, r *http.Request) {
	var food models.Food
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}

	err := services.CreateFood(&food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusCreated, food)
}

// Update a food
func UpdateFoodHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	var updatedFood models.Food
	if err := json.NewDecoder(r.Body).Decode(&updatedFood); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}

	err = services.UpdateFood(id, &updatedFood)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusAccepted, updatedFood)
}

// Delete a food
func DeleteFoodHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	err = services.DeleteFood(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
