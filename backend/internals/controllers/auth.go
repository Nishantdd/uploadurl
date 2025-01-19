package controllers

import (
	"log"
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/internals/service"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
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

func Logout(c *gin.Context) {
	c.SetCookie(config.Load().OAuth.SessionName, "", -1, "/", "", false, true)
	c.Redirect(302, config.Load().App.Address)
}
