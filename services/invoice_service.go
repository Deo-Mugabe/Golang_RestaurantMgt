package handlers

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
)

func GetInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	result := db.DB.Find(&invoices)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoices, nil
}

func GetInvoice(id uint) (*models.Invoice, error) {
	var invoice models.Invoice
	result := db.DB.First(&invoice, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &invoice, nil
}

func CreateInvoice(invoice *models.Invoice) error {
	result := db.DB.Create(&invoice)
	return result.Error

}

func UpdateInvoice(id uint, UpdatedInvoice *models.Invoice) error {
	var invoice models.Invoice
	result := db.DB.First(&invoice, id)
	if result.Error != nil {
		return result.Error
	}
	// update invoice

	db.DB.Save(&invoice)

	return nil
}

func DeleteInvoice(id uint) error {
	result := db.DB.Delete(&models.Invoice{}, id)
	return result.Error
}
