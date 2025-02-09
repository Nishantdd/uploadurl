package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/Nishantdd/uploadurl/backend/internals/middleware"
	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.RouterGroup) {
	fileGroup := r.Group("/file")
	fileGroup.Use(middleware.ValidateAuth())
	{
		fileGroup.GET("/user", controllers.GetFiles)
		fileGroup.GET("/", controllers.GetAllFiles)
		fileGroup.GET("/:id", controllers.GetFileByID)
		fileGroup.POST("/upload", middleware.StoreMultipartFilesLocally(), controllers.UploadFile)
		fileGroup.DELETE("/:id", controllers.DeleteFile)
	}
}
