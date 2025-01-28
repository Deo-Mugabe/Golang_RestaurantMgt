package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/services"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/utility"
)

func GetMenusHandler(w http.ResponseWriter, r *http.Request) {
	invoices, err := services.GetInvoices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, invoices)
}

func GetMenuHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	invoice, err := services.GetInvoice(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, invoice)
}

func CreateMenueHandler(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err := services.CreateInvoice(&invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusCreated, invoice)
}

func UpdateMenuHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	var updatedInvoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&updatedInvoice); err != nil {
		http.Error(w, "Invalid Inputs", http.StatusBadRequest)
		return
	}
	err = services.UpdateInvoice(uint(id), &updatedInvoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusAccepted, updatedInvoice)

}

func DeleteMenuHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	err = services.DeleteInvoice(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
