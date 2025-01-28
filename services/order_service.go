package services

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
)

func GetOrders() ([]models.Order, error) {
	var orders []models.Order
	result := db.DB.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func GetOrder(id uint) (*models.Order, error) {
	var order models.Order
	result := db.DB.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func CreateOrder(order *models.Order) error {
	result := db.DB.Create(&order)
	return result.Error
}

func UpdateOrder(id uint, UpdatedOrder *models.Order) error {
	var order models.Order
	result := db.DB.First(&order, id)
	if result.Error != nil {
		return result.Error
	}
	//update order

	db.DB.Save(&order)
	return nil
}

func DeleteOrder(id uint) error {
	result := db.DB.Delete(&models.Order{}, id)
	return result.Error
}
