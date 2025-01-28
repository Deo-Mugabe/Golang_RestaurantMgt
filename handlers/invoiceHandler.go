package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/services"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/utility"
)

// Get all invoices
func GetInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	invoices, err := services.GetInvoices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusOK, invoices)
}

// Get a single invoice by ID
func GetInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	invoice, err := services.GetInvoice(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	utility.JsonResponse(w, http.StatusOK, invoice)
}

// Create a new invoice
func CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := services.CreateInvoice(&invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusCreated, invoice)
}

// Update an existing invoice
func UpdateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedInvoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&updatedInvoice); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.UpdateInvoice(id, &updatedInvoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utility.JsonResponse(w, http.StatusAccepted, updatedInvoice)
}

// Delete an invoice
func DeleteInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utility.ParseID(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteInvoice(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
