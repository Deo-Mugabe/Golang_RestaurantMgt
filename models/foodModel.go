package models

import "time"

type Food struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string      `gorm:"type:varchar(100);not null" json:"name"`
	Price      float64     `gorm:"type:decimal(10,2);not null" json:"price"`
	FoodImage  string      `gorm:"type:varchar(255)" json:"food_image"` // Optional URL to image
	FoodID     string      `gorm:"type:varchar(50);unique;not null" json:"food_id"`
	MenuID     uint        `gorm:"not null" json:"menu_id"`       // Foreign key to Menu
	Menu       Menu        `gorm:"foreignKey:MenuID" json:"menu"` // Belongs to Menu
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	OrderItems []OrderItem `gorm:"foreignKey:FoodID" json:"order_items"` // One-to-Many relationship with OrderItems
}
