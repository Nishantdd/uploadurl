package main

import (
	"log"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/routes"
	"github.com/Nishantdd/uploadurl/backend/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db.SetupConnection()

	router := gin.Default()
	routes.HandleRoutes(router)

	if err := router.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("Listening on %v", cfg.Server.Address)
	}
}
