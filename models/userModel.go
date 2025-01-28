package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName    string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName     string    `gorm:"type:varchar(100);not null" json:"last_name"`
	Email        string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password     string    `gorm:"type:varchar(255);not null" json:"-"` // Hashed password
	Avatar       string    `gorm:"type:varchar(255)" json:"avatar"`     // Optional profile image URL
	Phone        string    `gorm:"type:varchar(15);unique" json:"phone"`
	Token        string    `gorm:"type:text" json:"token"`
	RefreshToken string    `gorm:"type:text" json:"refresh_token"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UserID       string    `gorm:"type:varchar(50);unique;not null" json:"user_id"`
}
