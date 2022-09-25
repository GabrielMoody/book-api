package models

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

var lock = &sync.Mutex{}

func GetConnection() *gorm.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		dsn := "root:@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err.Error())
			return nil
		}

		db.AutoMigrate(&Book{}, &User{})
	}

	return db
}
