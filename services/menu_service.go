package handlers

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
)

func GetMenus() ([]models.Menu, error) {
	var menus []models.Menu
	result := db.DB.Find(&menus)
	if result.Error != nil {
		return nil, result.Error
	}
	return menus, nil
}

func GetMenu(id uint) (*models.Menu, error) {
	var menu models.Menu
	result := db.DB.First(&menu, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &menu, nil
}

func CreateMenu(menu *models.Menu) error {
	result := db.DB.Create(&menu)
	return result.Error
}

func UpdateMenu(id uint, updatedMenu *models.Menu) error {
	var menu models.Menu
	result := db.DB.First(&menu, id)
	if result.Error != nil {
		return result.Error
	}
	//update menu

	db.DB.Save(&menu)
	return nil
}

func DeleteMenu(id uint) error {
	result := db.DB.Delete(&models.Menu{}, id)
	return result.Error
}
