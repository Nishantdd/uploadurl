package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/login", controllers.Login)
	router.GET("/signup", controllers.Signup)
	router.GET("/googlelogin", controllers.GoogleLogin)
	router.GET("/auth/callback", controllers.GoogleCallback)
}
