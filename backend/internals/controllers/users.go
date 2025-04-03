package controllers

import (
	"net/http"
	"time"

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

func GetUsername(c *gin.Context) {
	var user models.User
	id := c.GetUint64("userId")

	if err := database.DB.First(&user, "ID = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username})
}

func GetUserMetadata(c *gin.Context) {
	var user models.User
	id := c.GetUint64("userId")

	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": id})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username":              user.Username,
		"username_permission":   time.Since(user.UpdatedAt) >= 30*24*time.Hour,
		"repository_permission": true, // TODO: To be implemented
	})
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

func UpdateUsername(c *gin.Context) {
	var user models.User
	var usernameReq struct {
		Username string `json:"username" binding:"required,min=1,max=100"`
	}
	id := c.GetUint64("userId")

	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&usernameReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.First(&user, "username = ?", usernameReq.Username).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	user.Username = usernameReq.Username
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Username updated successfully"})
}

func UpdatePassword(c *gin.Context) {
	var user models.User
	var passwordReq models.PasswordRequest
	id := c.GetUint64("userId")

	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&passwordReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CompareHash(user.Password, passwordReq.OldPassword); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Old password doesn't match"})
		return
	}

	user.Password = utils.Hash(passwordReq.NewPassword)
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
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
	userId := c.GetUint64("userId")

	result := database.DB.Delete(&models.User{}, userId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func GetUserAuth(c *gin.Context) {
	var user models.User
	userId := c.GetUint64("userId")

	if err := database.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUserAuth(c *gin.Context) {
	var user models.User
	var userReq models.UserRequest
	userId := c.GetUint64("userId")

	// Check if user exists
	if err := database.DB.First(&user, userId).Error; err != nil {
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

func DeleteUserAuth(c *gin.Context) {
	userId := c.GetUint64("userId")

	result := database.DB.Delete(&models.User{}, userId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
