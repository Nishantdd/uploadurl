package main

import (
	"log"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/routes"
	"github.com/Nishantdd/uploadurl/backend/internals/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	database.DB = database.SetupConnection()

	service.Oauth2Config, service.Oauth2State = service.InitOauth()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	routes.HandleRoutes(router)

	if err := router.Run(cfg.Server.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("Listening on %v", cfg.Server.ServerAddress)
	}
}
