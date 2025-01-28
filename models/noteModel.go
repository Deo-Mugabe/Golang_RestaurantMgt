package models

import "time"

type Note struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Text      string    `gorm:"type:text;not null" json:"text"`
	NoteID    string    `gorm:"type:varchar(50);unique;not null" json:"note_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
