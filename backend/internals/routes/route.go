package routes

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	UserRoutes(apiGroup)
	AuthRoutes(apiGroup)
}
