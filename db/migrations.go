package database

import (
	"golang-prototype/models"
	"log"

	"gorm.io/gorm"
)

// MigrateDB handles all database migrations
func MigrateDB(db *gorm.DB) {
	// Auto-migrate your models
	err := db.AutoMigrate(
		&models.User{}, // Add more models here as your project grows
	)
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	log.Println("✅ Database migrated successfully")
}

