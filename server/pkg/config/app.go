package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectSqlite() {
	d, err := gorm.Open(sqlite.Open("easylease.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
