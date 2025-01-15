package main

import (
	"log"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	database.DB = database.SetupConnection()

	router := gin.Default()
	routes.HandleRoutes(router)

	if err := router.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("Listening on %v", cfg.Server.Address)
	}
}
