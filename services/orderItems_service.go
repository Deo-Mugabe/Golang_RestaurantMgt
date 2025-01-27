package handlers

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
)

func GetOrderItemsByOrder() ([]models.OrderItem, error) {
	return nil, nil
}

func GetOrderItems() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	result := db.DB.Find(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderItems, nil
}

func GetOrderItem(id uint) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	result := db.DB.First(&orderItem, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orderItem, nil
}

func CreateOrderItem(orderItem *models.OrderItem) error {
	result := db.DB.Create(&orderItem)
	return result.Error
}

func UpdateOrderItem(id uint, updatedOrderItem *models.OrderItem) error {
	var orderItem models.OrderItem
	result := db.DB.First(&orderItem, id)
	if result.Error != nil {
		return result.Error
	}
	//update order

	db.DB.Save(&orderItem)
	return nil
}

func DeleteOrderItem(id uint) error {
	result := db.DB.Delete(&models.OrderItem{}, id)
	return result.Error
}
