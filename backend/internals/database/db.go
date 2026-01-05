package database

import (
	"log"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetupConnection() *gorm.DB {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Postgres.URI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalln(err)
	}

	// migrating Models/Schemas
	if err := db.AutoMigrate(&models.User{}, &models.Url{}, &models.Slug{}, &models.File{}, &models.Token{}, &models.UrlHits{}); err != nil {
		log.Fatalln("Failed to migrate database:", err)
	}

	return db
}
