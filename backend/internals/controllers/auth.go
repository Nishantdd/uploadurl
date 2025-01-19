package controllers

import (
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/Nishantdd/uploadurl/backend/internals/service"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func HandleGoogleLogin(c *gin.Context) {
	url := service.Oauth2Config.AuthCodeURL(service.Oauth2State, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func HandleGoogleCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	// Exchange the authorization code for an access token
	exchangeToken, err := service.Oauth2Config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to exchange Token"})
		return
	}

	// Use the access Token to get the user's Google profile
	client := service.Oauth2Config.Client(c, exchangeToken)
	userInfo, err := service.GetGoogleUserInfo(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user info"})
		return
	}

	// Creating User if not exists
	var user models.User
	userRes := database.DB.Where("email = ?", userInfo.Email).First(&user)
	if userRes.Error != nil {
		user = models.User{
			Username: userInfo.Email,
			Email:    userInfo.Email,
			Fullname: userInfo.Name,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// Hashing Token
	hashToken := utils.Hash(userInfo.Email, userInfo.Name)

	// Creating Token if not exists
	var token models.Token
	tokenRes := database.DB.Where("token = ?", hashToken).First(&token)
	if tokenRes.Error != nil {
		token = models.Token{
			Token: hashToken,
		}
		if err := database.DB.Create(&token).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"avatar": userInfo.Picture,
		"token":  token,
		"user":   user,
	})
}
