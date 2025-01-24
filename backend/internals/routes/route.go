package routes

import (
	"github.com/Nishantdd/uploadurl/backend/internals/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(router *gin.Engine) {
	// unprotected routes
	AuthRoutes(router)

	Group := router.Group("/api")
	UserRoutes(Group)

	// optionally protected routes (works with & without authToken)
	optionalGroup := router.Group("/api")
	optionalGroup.Use(middleware.ValidateOptionalAuth())
	{
		UrlRoutes(optionalGroup)
	}

	// protected routes (works only with authToken)
	protectedGroup := router.Group("/api")
	protectedGroup.Use(middleware.ValidateAuth())
	{
		UserRoutesAuth(protectedGroup)
	}
}
