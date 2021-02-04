package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}
