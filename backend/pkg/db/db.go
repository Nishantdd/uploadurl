package db

import (
	"log"

	"github.com/Nishantdd/uploadurl/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Postgres.URI), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
