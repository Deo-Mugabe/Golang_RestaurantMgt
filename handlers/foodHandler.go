package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/services"
	"github.com/gorilla/mux"
)

func GetFoodsHandler(w http.ResponseWriter, r *http.Request) {
	foods, err := services.GetFoods()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

func GetFoodHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	food, err := services.GetFood(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(food)
}

func CreateFoodHandler(w http.ResponseWriter, r *http.Request) {
	var food models.Food
	if err := json.NewDecoder(r.Body).Decode(food); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err := services.CreateFood(&food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(food)
}

func UpdateFoodHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	var updatedFood models.Food
	if err := json.NewDecoder(r.Body).Decode(updatedFood); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err = services.UpdateFood(uint(id), &updatedFood)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(updatedFood)
}

func DeleteFoodHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	err = services.DeleteFood(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
