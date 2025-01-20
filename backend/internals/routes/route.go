package routes

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.Engine) {
	AuthRoutes(router)

	apiGroup := router.Group("/api")
	UserRoutes(apiGroup)
	UrlRoutes(apiGroup)
}
