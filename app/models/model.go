package models

import "gorm.io/gorm"

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Book{})
	return db
}
