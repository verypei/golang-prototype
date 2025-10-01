package database

import (
	"fmt"
	"log"

	"golang-prototype/config"
	"golang-prototype/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := config.Config.DBUrl // ✅ use Config.DBUrl instead of config.DB_URL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db

	// Auto migrate tables
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("✅ Database connected successfully")
}
