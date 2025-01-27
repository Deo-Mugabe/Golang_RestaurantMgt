package services

import (
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/db"
	"github.com/Deo-Mugabe/Golang_RestaurantMgt/Golang_RestaurantMgt/models"
)

func GetTables() ([]models.Table, error) {
	var tables []models.Table
	result := db.DB.Find(&tables)
	if result.Error != nil {
		return nil, result.Error
	}
	return tables, nil
}

func GetTable(id uint) (*models.Table, error) {
	var table models.Table
	result := db.DB.First(&table, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &table, nil
}

func CreateTable(table *models.Table) error {
	result := db.DB.Create(&table)
	return result.Error
}

func UpdateTable(id uint, updatedTable *models.Table) error {
	var table models.Table
	result := db.DB.First(&table, id)
	if result.Error != nil {
		return result.Error
	}
	//update order

	db.DB.Save(&table)
	return nil
}

func DeleteTable(id uint) error {
	result := db.DB.Delete(&models.Table{}, id)
	return result.Error
}
