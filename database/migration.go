package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./mygram.db"), &gorm.Config{})
	if err == nil {
		return db, err
	}
	return db, nil
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	log.Println("Migration success")
}