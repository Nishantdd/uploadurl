package controllers

import (
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/gin-gonic/gin"
)

func GetFiles(c *gin.Context) {
	var files []models.File
	userId, _ := c.Get("userId")

	if err := database.DB.Find(&files, "userId = ?", userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func GetFileByID(c *gin.Context) {
	var file models.File
	fileId := c.Param("id")

	if err := database.DB.First(&file, "ID = ?", fileId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file": file})
}

func UploadFile(c *gin.Context) {
	// TODO: Needs s3 wrapper for implementation
}

func DeleteFile(c *gin.Context) {
	fileId := c.Param("id")

	result := database.DB.Delete(&models.File{}, "ID = ?", fileId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file deleted successfully"})
}
