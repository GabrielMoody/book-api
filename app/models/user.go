package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null;email"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUserByUsername(db *gorm.DB, username string) *User {
	var user *User

	result := db.Where("username = ?", username).Find(&user)

	if result.RowsAffected < 1 {
		return nil
	}

	return user
}
