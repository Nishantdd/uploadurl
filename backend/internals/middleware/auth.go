package middleware

import (
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/gin-gonic/gin"
)

func ValidateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if token exists
		var token models.Token
		tokenRes := database.DB.Where("token = ?", authToken).First(&token)
		if tokenRes.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised Access not allowed!!!"})
			return
		}

		// Get user from database
		var user models.User
		if err := database.DB.First(&user, token.UserId).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Set user in context
		c.Set("userId", user.ID)
		c.Next()
	}
}
