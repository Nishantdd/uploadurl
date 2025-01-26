package routes

import (
	"github.com/gin-gonic/gin"
)

func HandleRoutes(router *gin.Engine) {
	// unprotected routes
	AuthRoutes(router)

	Group := router.Group("/api")
	UserRoutes(Group)
	UrlRoutes(Group)
	UserRoutesAuth(Group)
	FileRoutes(Group)
}
