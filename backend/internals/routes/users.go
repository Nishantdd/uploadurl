package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("", controllers.GetUsers)          // read all
		userGroup.POST("", controllers.CreateUser)       // create
		userGroup.GET("/:id", controllers.GetUserByID)   // read one
		userGroup.PUT("/:id", controllers.UpdateUser)    // update
		userGroup.DELETE("/:id", controllers.DeleteUser) // delete
	}
}
