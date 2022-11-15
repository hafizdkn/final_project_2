package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = os.Getenv("DB_HOST")
	port   = os.Getenv("DB_PORT")
	user   = os.Getenv("DB_USER")
	pass   = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB_NAME")
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Photo{}, &Comment{}, &SocialMedia{})

	return db, nil
}
