package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("", controller.GetUsers)          // read all
		userGroup.POST("", controller.CreateUser)       // create
		userGroup.GET("/:id", controller.GetUserByID)   // read one
		userGroup.PUT("/:id", controller.UpdateUser)    // update
		userGroup.DELETE("/:id", controller.DeleteUser) // delete
	}
}
