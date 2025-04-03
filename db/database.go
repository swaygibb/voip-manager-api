package db

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"voip-manager-api/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("voip_manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.CallLog{})
	DB.AutoMigrate(&models.Voicemail{})
}
