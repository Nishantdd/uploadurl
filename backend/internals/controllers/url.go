package controllers

import (
	"net/http"
	"net/url"

	"github.com/Nishantdd/uploadurl/backend/config"
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
	userId, exists := c.Get("userId")

	if err := c.ShouldBindJSON(&urlReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := url.Parse(urlReq.Url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	slugValue := utils.GenerateUniqueString(ShortURLLength)

	if exists {
		userIdValue := userId.(uint64)
		if err := service.RegisterUrl(urlReq.Url, slugValue, "short", &userIdValue); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := service.RegisterUrl(urlReq.Url, slugValue, "short", nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url": config.Load().Server.DomainAddress + "/" + slugValue,
		"url":       urlReq.Url,
	})
}
