package models

import "time"

type OrderItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Quantity  int       `gorm:"type:int;not null" json:"quantity"`
	UnitPrice float64   `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	FoodID    uint      `gorm:"not null" json:"food_id"`  // Foreign key referencing Food
	OrderID   uint      `gorm:"not null" json:"order_id"` // Foreign key referencing Order
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
