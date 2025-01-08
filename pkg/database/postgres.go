package database

import (
	"github.com/devenock/romax-backend/config"
	"github.com/devenock/romax-backend/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	err = DB.AutoMigrate(&models.Player{}, &models.Game{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
