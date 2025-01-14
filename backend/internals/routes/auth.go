package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", controllers.Signup) // Signup (create a new user)
		authGroup.POST("/login", controllers.Login)   // Login (authenticate user)
	}
}
