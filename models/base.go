package models

import (
	"github.com/jinzhu/gorm"

	"github.com/msinwolc/config"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func OpenDB() {
	db = config.InitDB()
}

func CloseDB() {
	db.Close()
}
