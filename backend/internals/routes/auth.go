package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	{
		authGroup.GET("/:provider/callback", controllers.HandleAuthCallback)
		authGroup.GET("/logout/:provider", controllers.HandleLogout)
		authGroup.GET("/:provider", controllers.HandleProviderLogin)
	}
}
