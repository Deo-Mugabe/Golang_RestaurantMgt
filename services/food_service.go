package services

import (
	"errors"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/google/uuid"
)

// Get all foods from the database
func GetFoods() ([]models.Food, error) {
	var foods []models.Food
	result := db.DB.Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}
	return foods, nil
}

// Get a single food item by ID
func GetFood(id uint) (*models.Food, error) {
	var food models.Food
	result := db.DB.First(&food, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &food, nil
}

// Create a new food item in the database
func CreateFood(food *models.Food) error {
	// Ensure the food object is valid
	if food.Name == "" || food.Price <= 0 {
		return errors.New("invalid food data: name and price are required")
	}

	// Check if the associated Menu exists
	var menu models.Menu
	result := db.DB.First(&menu, food.MenuID)
	if result.Error != nil {
		return errors.New("menu not found")
	}

	// Ensure `food_id` is unique (generate if not provided)
	if food.FoodID == "" {
		food.FoodID = uuid.New().String() // Generate unique food ID
	}

	// Save the new food item
	result = db.DB.Create(food)
	return result.Error
}

// Update an existing food item
func UpdateFood(id uint, updatedFood *models.Food) error {
	var food models.Food
	result := db.DB.First(&food, id)
	if result.Error != nil {
		return errors.New("food item not found")
	}

	// Ensure valid data before updating
	if updatedFood.Name == "" || updatedFood.Price <= 0 {
		return errors.New("invalid update: name and price are required")
	}

	// Prevent modification of FoodID
	updatedFood.FoodID = food.FoodID

	// Update food properties
	food.Name = updatedFood.Name
	food.Price = updatedFood.Price
	food.FoodImage = updatedFood.FoodImage

	// Save the updated food item
	result = db.DB.Save(&food)
	return result.Error
}

// Delete a food item by ID
func DeleteFood(id uint) error {
	result := db.DB.Delete(&models.Food{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func round(num float64) int {
// 	return 0
// }

// func toFixed(num float64, precision int) float64 {
// 	return 0
// }
