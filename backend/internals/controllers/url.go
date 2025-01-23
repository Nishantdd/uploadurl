package controllers

import (
	"net/http"
	"net/url"

	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/Nishantdd/uploadurl/backend/internals/service"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
	"github.com/gin-gonic/gin"
)

const (
	ShortURLLength = 5
)

func ShortenUrl(c *gin.Context) {
	var urlReq models.UrlRequest
	userId, _ := c.Get("userId")

	if err := c.ShouldBindJSON(&urlReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := url.Parse(urlReq.Url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
	}

	slugValue := utils.GenerateUniqueString(ShortURLLength)

	if userId != nil {
		if err := service.RegisterUrl(urlReq.Url, slugValue, "short", userId.(*uint64)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		if err := service.RegisterUrl(urlReq.Url, slugValue, "short", nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
