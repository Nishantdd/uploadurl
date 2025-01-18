package controllers

import (
	"log"
	"net/http"

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
	token, err := service.Oauth2Config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to exchange token"})
		return
	}

	// Use the access token to get the user's Google profile
	client := service.Oauth2Config.Client(c, token)
	userInfo, err := service.GetGoogleUserInfo(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user info"})
		return
	}

	tempToken := utils.Hash(userInfo.Email, userInfo.Name)
	log.Println(tempToken)
	c.JSON(http.StatusOK, gin.H{
		"avatar": userInfo.Picture,
		"token":  utils.Hash(userInfo.Email, userInfo.Name),
	})
}
