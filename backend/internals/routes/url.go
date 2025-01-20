package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/controllers"
	"github.com/gin-gonic/gin"
)

func UrlRoutes(r *gin.RouterGroup) {
	urlGroup := r.Group("/url")
	{
		urlGroup.POST("/shorten", controllers.ShortenUrl)
	}
}
