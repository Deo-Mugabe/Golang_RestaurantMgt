package models

import "time"

type Invoice struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	InvoiceID      string    `gorm:"type:varchar(50);unique;not null" json:"invoice_id"`
	OrderID        uint      `gorm:"not null" json:"order_id"`                        // Foreign key to Order
	Order          Order     `gorm:"foreignKey:OrderID" json:"order"`                 // Belongs to Order
	PaymentMethod  string    `gorm:"type:varchar(50);not null" json:"payment_method"` // CARD, CASH, etc.
	PaymentStatus  string    `gorm:"type:varchar(20);not null" json:"payment_status"` // PAID, PENDING, etc.
	PaymentDueDate time.Time `gorm:"type:date;not null" json:"payment_due_date"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
