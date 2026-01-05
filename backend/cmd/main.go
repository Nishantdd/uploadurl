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
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	routes.HandleRoutes(router)

	log.Fatal(router.Run(cfg.Server.ServerAddress))
}
