package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() error {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, 5432)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}
	fmt.Println("database connected!")

	DB = db
	return nil
}
