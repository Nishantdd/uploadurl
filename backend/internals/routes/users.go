package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/Nishantdd/uploadurl/backend/internals/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.DELETE("/", middleware.ValidateAuth(), controllers.DeleteUser)
		userGroup.PATCH("/password", middleware.ValidateAuth(), controllers.UpdatePassword)
		userGroup.PATCH("/username", middleware.ValidateAuth(), controllers.UpdateUsername)
		userGroup.GET("/username", middleware.ValidateAuth(), controllers.GetUsername)
		userGroup.GET("/metadata", middleware.ValidateAuth(), controllers.GetUserMetadata)
	}
}

func UserRoutesAuth(r *gin.RouterGroup) {
	authGroup := r.Group("/auth/profile")
	authGroup.Use(middleware.ValidateAuth())
	{
		authGroup.GET("/", controllers.GetUserAuth)
		authGroup.PUT("/", controllers.UpdateUserAuth)
		authGroup.DELETE("/", controllers.DeleteUserAuth)
	}
}
