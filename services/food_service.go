package services

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
)

func GetFoods() ([]models.Food, error) {
	var foods []models.Food
	result := db.DB.Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}
	return foods, nil
}

func GetFood(id uint) (*models.Food, error) {
	var food models.Food
	result := db.DB.First(&food, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &food, nil
}

func CreateFood(food *models.Food) error {
	result := db.DB.Create(&food)

	return result.Error
}

func UpdateFood(id uint, updatedFood *models.Food) error {
	var food models.Food
	result := db.DB.First(&food, id)
	if result.Error != nil {
		return result.Error
	}
	//update food
	db.DB.Save(&food)
	return nil
}

func DeleteFood(id uint) error {
	result := db.DB.Delete(&models.Food{}, id)
	return result.Error
}
