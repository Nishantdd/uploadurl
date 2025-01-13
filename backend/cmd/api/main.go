package main

import (
	"log"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	if err := r.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("Listening on %v", cfg.Server.Address)
	}
}
