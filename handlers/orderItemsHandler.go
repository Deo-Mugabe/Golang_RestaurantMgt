package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/services"
	"github.com/gorilla/mux"
)

func GetOrderItemsHandler(w http.ResponseWriter, r *http.Request) {
	orderItems, err := services.GetOrderItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderItems)
}

func GetOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	orderItem, err := services.GetOrderItem(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderItem)
}

func GetOrderItemsByOrderHandler(w http.ResponseWriter, r *http.Request) {
	var orderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(orderItem); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err := services.CreateOrderItem(&orderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(orderItem)
}

func CreateOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	var orderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(orderItem); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err := services.CreateOrderItem(&orderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(orderItem)
}

func UpdateOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	var updatedOrderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(updatedOrderItem); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err = services.UpdateOrderItem(uint(id), &updatedOrderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(updatedOrderItem)
}

func DeleteOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	err = services.DeleteOrderItem(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
