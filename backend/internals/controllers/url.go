package controllers

import (
	"net/http"
	"net/url"
	"time"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/database"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide an url"})
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

func GetAllUrls(c *gin.Context) {
	var urls []models.Url
	urlRes := database.DB.Find(&urls)
	if urlRes.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": urlRes.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"urls": urls})
}

func GetUrls(c *gin.Context) {
	userId, _ := c.Get("userId")

	var urls []models.Url
	urlRes := database.DB.Find(&urls, "user_id = ?", userId)
	if urlRes.Error != nil || len(urls) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Url Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"urls": urls})
}

func DeleteUrl(c *gin.Context) {
	id := c.Param("id")

	var url models.Url
	urlRes := database.DB.Where("id = ?", id).First(&url)
	if urlRes.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Url Found"})
		c.Abort()
		return
	}

	var slug models.Slug
	slugRes := database.DB.Where("slug = ?", url.Slug).First(&slug)
	if slugRes.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Slug Found"})
		c.Abort()
		return
	}

	delSlugRes := database.DB.Delete(&models.Slug{}, slug.ID)
	if delSlugRes.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": delSlugRes.Error.Error()})
		return
	}

	delUrlRes := database.DB.Delete(&models.Url{}, url.ID)
	if delUrlRes.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": delUrlRes.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL deleted successfully"})
}

func UpdateUrlHits(c *gin.Context) {
	slug := c.Param("slug")

	var url models.Url
	urlRes := database.DB.Where("slug = ?", slug).First(&url)
	if urlRes.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Slug Found"})
		c.Abort()
		return
	}

	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	var HitsObj models.UrlHits
	if err := database.DB.Where("url_id = ? AND slug = ? AND created_at BETWEEN ? AND ?", url.ID, slug, startOfDay, endOfDay).First(&HitsObj).Error; err != nil {
		// Create new hit record if not found
		HitsObj = models.UrlHits{
			UrlId: url.ID,
			Slug:  slug,
			Hits:  1,
		}
		if err := database.DB.Create(&HitsObj).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hit record"})
			return
		}
	} else {
		// Update existing hit record
		HitsObj.Hits++
		if err := database.DB.Save(&HitsObj).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hit record"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hit count updated"})
}
