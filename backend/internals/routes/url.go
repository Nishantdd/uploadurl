package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/Nishantdd/uploadurl/backend/internals/middleware"
	"github.com/gin-gonic/gin"
)

func UrlRoutes(r *gin.RouterGroup) {
	urlGroup := r.Group("/url")
	{
		urlGroup.GET("/user", middleware.ValidateAuth(), controllers.GetUrls)
		urlGroup.GET("/", middleware.ValidateAuth(), controllers.GetAllUrls)
		urlGroup.POST("/shorten", middleware.ValidateOptionalAuth(), controllers.ShortenUrl)
		urlGroup.PATCH("/:slug", controllers.UpdateUrlHits)
		urlGroup.DELETE("/:id", middleware.ValidateAuth(), controllers.DeleteUrl)
	}
}
