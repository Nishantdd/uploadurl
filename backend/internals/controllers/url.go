package controllers

import (
	"net/http"
	"net/url"

	"github.com/Nishantdd/uploadurl/backend/internals/service"
	"github.com/gin-gonic/gin"
)

const (
	ShortURLLength = 5
)

func ShortenUrl(c *gin.Context) {
	var body struct {
		Url    string  `json:"url" binding:"required"` // Original url of the user
		UserId *uint64 `json:"user_id"`                // UserId (optional)
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := url.Parse(body.Url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
	}

	slugValue := service.GenerateRandomSlug(ShortURLLength)
	if err := service.RegisterUrl(body.Url, slugValue, "short", nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
