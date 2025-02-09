package controllers

import (
	"net/http"
	"os"
	"path"

	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/Nishantdd/uploadurl/backend/internals/service"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
	"github.com/gin-gonic/gin"
)

func GetAllFiles(c *gin.Context) {
	var files []models.File
	if err := database.DB.Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func GetFiles(c *gin.Context) {
	var files []models.File
	userId, _ := c.Get("userId")

	if err := database.DB.Find(&files, "user_id = ?", userId).Error; err != nil {
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
	filepath, exists := c.Get("path")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't find the uploaded file on server"})
		return
	}

	file, err := os.Open(filepath.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file: " + err.Error()})
		return
	}
	defer file.Close()

	s3Client, err := service.NewS3Client()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileHash := utils.GenerateUniqueString(16)
	if location, err := s3Client.UploadFile(file, fileHash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		fileName := path.Base(file.Name())
		fileStats, _ := file.Stat()
		fileExt := path.Ext(file.Name())

		res := database.DB.Create(&models.File{
			FileName: fileName,
			FileHash: fileHash,
			FileType: fileExt,
			FileSize: fileStats.Size(),
			Location: location,
			UserId:   userId.(uint64),
		})
		if res.Error != nil {
			s3Client.DeleteFile(fileHash)
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"location": location})
	}
}

func DeleteFile(c *gin.Context) {
	fileId := c.Param("id")
	var file models.File
	result := database.DB.Find(&file, "ID = ?", fileId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
	}

	s3Client, err := service.NewS3Client()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	s3Client.DeleteFile(file.FileHash)

	result = database.DB.Delete(&file)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file deleted successfully"})
}
