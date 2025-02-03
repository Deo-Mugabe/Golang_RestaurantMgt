package models

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderDate  time.Time   `gorm:"type:datetime;not null" json:"order_date"`
	TableID    uint        `gorm:"not null" json:"table_id"` // Foreign key to Table
	OrderID    string      `gorm:"type:varchar(50);unique;not null" json:"order_id"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"` // One-to-Many relationship with OrderItems
}
