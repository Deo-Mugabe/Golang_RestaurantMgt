package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/services"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/utility"
)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := services.GetOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, orders)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	order, err := services.GetOrder(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, order)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err := services.CreateOrder(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusCreated, order)
}

func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	var updatedOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err = services.UpdateOrder(uint(id), &updatedOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusAccepted, updatedOrder)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	err = services.DeleteOrder(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
