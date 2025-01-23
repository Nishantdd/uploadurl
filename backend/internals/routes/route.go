package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(router *gin.Engine) {
	AuthRoutes(router)

	Group := router.Group("")
	UserRoutes(Group)

	// protected routes
	protectedGroup := router.Group("/api")
	protectedGroup.Use(middleware.ValidateAuth())
	{
		UrlRoutes(protectedGroup)
	}
}
