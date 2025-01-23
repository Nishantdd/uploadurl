package controllers

import (
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateUser(c *gin.Context) {
	var userReq models.UserRequest
	var user models.User

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Username = userReq.Username
	user.Email = userReq.Email
	user.Fullname = userReq.Fullname
	user.Avatar = userReq.Avatar

	// Hashing password
	user.Password = utils.Hash(userReq.Password)

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	var userReq models.UserRequest
	id := c.Param("id")

	// Check if user exists
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind update data
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Email = userReq.Email
	user.Username = userReq.Username
	user.Fullname = userReq.Fullname
	user.Avatar = userReq.Avatar

	// Hashing password
	user.Password = utils.Hash(userReq.Password)

	// Update user
	result := database.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
