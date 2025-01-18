package controllers

import (
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/gin-gonic/gin"
)

func GoogleCallback(c *gin.Context) {
	c.Redirect(302, config.Load().App.Address)
}

func UserInfoHandler(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User authenticated",
		"user":    user,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(config.Load().OAuth.SessionName, "", -1, "/", "", false, true)
	c.Redirect(302, config.Load().App.Address)
}
