package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/login", controllers.HandleGoogleLogin)
	router.GET("/auth/callback", controllers.HandleGoogleCallback)
}
