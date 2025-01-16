package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/gin-gonic/gin"
	"github.com/zalando/gin-oauth2/google"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/login", google.LoginHandler)
	router.GET("/logout", controllers.Logout)

	AuthGroup := router.Group("/auth")
	AuthGroup.Use(google.Auth())
	{
		AuthGroup.GET("/user", controllers.UserInfoHandler)
	}
}
