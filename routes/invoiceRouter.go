package routes

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/handlers"
	"github.com/gorilla/mux"
)

func InvoiceRouter(r *mux.Router) {
	r.HandleFunc("/invoices", handlers.GetInvoicesHandler).Methods("GET")
	r.HandleFunc("/invoices/{id:[0-9]+}", handlers.GetInvoiceHandler).Methods("GET")
	r.HandleFunc("/invoices", handlers.CreateInvoiceHandler).Methods("POST")
	r.HandleFunc("/invoices/{id:[0-9]+}", handlers.UpdateInvoiceHandler).Methods("PUT")
	r.HandleFunc("/invoices/{id:[0-9]+}", handlers.DeleteInvoiceHandler).Methods("DELETE")
}
