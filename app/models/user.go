package models

import (
	"time"
)

type User struct {
	ID        uint
	Username  string `gorm:"unique not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var conn = GetConnection()

func GetUserByUsername(username string) *User {
	var user *User

	conn.Where("username = ?", username).Find(&user)

	return user
}
