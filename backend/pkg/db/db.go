package db

import (
	"log"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupConnection() {
	cfg := config.Load()

	DB, err := gorm.Open(postgres.Open(cfg.Postgres.URI), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// migrating Models/Schemas
	if err := DB.AutoMigrate(&models.User{}, &models.Url{}, &models.Slug{}, &models.Files{}); err != nil {
		log.Fatalln("Failed to migrate database:", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
