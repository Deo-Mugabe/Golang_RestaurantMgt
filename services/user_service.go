package services

import (
	"errors"
	"time"

	"github.com/Deo-Mugabe/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/models"
	"github.com/google/uuid"
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
	// Ensure required fields are present
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required fields")
	}

	// Hash the password before storing it
	err := user.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Set timestamps correctly
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Generate a unique UserID
	user.UserID = uuid.New().String()

	// Save user to database
	result := db.DB.Create(user)
	return result.Error
}

func UpdateUser(id uint, updatedUser *models.User) error {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	// Update fields only if they are provided
	if updatedUser.FirstName != "" {
		user.FirstName = updatedUser.FirstName
	}
	if updatedUser.LastName != "" {
		user.LastName = updatedUser.LastName
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Avatar != "" {
		user.Avatar = updatedUser.Avatar
	}
	if updatedUser.Phone != "" {
		user.Phone = updatedUser.Phone
	}
	if updatedUser.Password != "" {
		// Hash new password before saving
		err := user.HashPassword(updatedUser.Password)
		if err != nil {
			return err
		}
	}

	// Update timestamp
	user.UpdatedAt = time.Now()

	// Save the updated user
	result = db.DB.Save(&user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := db.DB.Delete(&models.User{}, id)

	return result.Error
}
