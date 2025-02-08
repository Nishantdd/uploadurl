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
		userGroup.GET("/username", middleware.ValidateAuth(), controllers.GetUsername)
		userGroup.POST("", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
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
