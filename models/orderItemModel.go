package models

import "time"

type OrderItem struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Quantity    int       `gorm:"type:int;not null" json:"quantity"`
	UnitPrice   float64   `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	OrderID     uint      `gorm:"not null" json:"order_id"`        // Foreign key to Order
	Order       Order     `gorm:"foreignKey:OrderID" json:"order"` // Belongs to Order
	FoodID      uint      `gorm:"not null" json:"food_id"`         // Foreign key to Food
	Food        Food      `gorm:"foreignKey:FoodID" json:"food"`   // Belongs to Food
	OrderItemID string    `gorm:"type:varchar(50);unique;not null" json:"order_item_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
