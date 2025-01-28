package models

import "time"

type Menu struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Category  string    `gorm:"type:varchar(50);not null" json:"category"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;not null" json:"end_date"`
	MenuID    string    `gorm:"type:varchar(50);unique;not null" json:"menu_id"`
	Foods     []Food    `gorm:"foreignKey:MenuID" json:"foods"` // One-to-Many relationship with Foods
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
