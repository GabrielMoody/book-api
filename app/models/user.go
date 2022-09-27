package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Username  string `gorm:"unique not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUserByUsername(db *gorm.DB, username string) *User {
	var user *User

	db.Where("username = ?", username).Find(&user)

	return user
}
