package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/services"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/utility"
)

func GetOrderItemsHandler(w http.ResponseWriter, r *http.Request) {
	orderItems, err := services.GetOrderItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, orderItems)
}

func GetOrderItemHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	orderItem, err := services.GetOrderItem(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, orderItem)
}

// GetOrderItemsByOrderHandler fetches all order items for a specific order
func GetOrderItemsByOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Extract order_id from the URL

	orderID, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	// Call the service to get order items
	orderItems, err := services.GetOrderItemsByOrder(uint(orderID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the order items in JSON format
	utility.JsonResponse(w, http.StatusOK, orderItems)
}

func CreateOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	var orderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err := services.CreateOrderItem(&orderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusCreated, orderItem)
}

func UpdateOrderItemHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	var updatedOrderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&updatedOrderItem); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err = services.UpdateOrderItem(uint(id), &updatedOrderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusAccepted, updatedOrderItem)
}

func DeleteOrderItemHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
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
