package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

var cfg *config.Config = config.Load()

func HandleProviderLogin(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))

	// try to get the user without re-authenticating
	if user, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Println(user)
		c.Redirect(http.StatusOK, cfg.App.Address)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func HandleAuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(user)
	c.Redirect(http.StatusOK, cfg.App.Address)
}

func HandleLogout(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusOK, cfg.App.Address)
}
