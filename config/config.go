package config

import (
	"fmt"
	"assignment-2-golang-hacktiv8/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
	host		= "localhost"
	user		= "postgres"
	password	= "root"
	dbPort		= "5432"
	dbName		= "go-assignment-2"
	db			*gorm.DB
	err			error
)

func StartDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	defer fmt.Println("Successfully Connected to Database")

	db.AutoMigrate(models.Order{}, models.Item{})
	return db
}