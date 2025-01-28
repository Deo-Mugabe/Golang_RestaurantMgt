package models

import "time"

type Table struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TableNumber    string    `gorm:"type:varchar(20);not null;unique" json:"table_number"`
	NumberOfGuests int       `gorm:"type:int;not null" json:"number_of_guests"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	TableID        string    `gorm:"type:varchar(50);unique;not null" json:"table_id"`
	Orders         []Order   `gorm:"foreignKey:TableID" json:"orders"` // One-to-Many relationship with Orders
}
