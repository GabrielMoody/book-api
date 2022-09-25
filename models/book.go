package models

import "time"

type Book struct {
	ID          uint
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	Rating      int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
