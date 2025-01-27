package services

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUser(id uint) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	result := db.DB.Create(&user)
	return result.Error
}

func UpdateUser(id uint, updatedUser *models.User) error {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	//update order

	db.DB.Save(&user)
	return nil
}

func DeleteUser(id uint) error {
	result := db.DB.Delete(&models.User{}, id)
	return result.Error
}
