package controllers

import (
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/Nishantdd/uploadurl/backend/internals/service"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

const (
	usernameLength = 4
)

func GoogleLogin(c *gin.Context) {
	url := service.Oauth2Config.AuthCodeURL(service.Oauth2State, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	// Exchange the authorization code for an access token
	exchangeToken, err := service.Oauth2Config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to exchange Token"})
		return
	}

	// Use the access Token to get the user's Google profile
	client := service.Oauth2Config.Client(c, exchangeToken)
	userInfo, err := service.GetGoogleUserInfo(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user info"})
		return
	}

	// Generating unique username of length
	username := utils.GenerateUniqueString(usernameLength)

	// Creating User if not exists
	var user models.User
	userRes := database.DB.Where("email = ?", userInfo.Email).First(&user)
	if userRes.Error != nil {
		user = models.User{
			Username: username,
			Email:    userInfo.Email,
			Fullname: userInfo.Name,
			Avatar:   userInfo.Avatar,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// Hashing Token
	hashToken := utils.Hash(userInfo.Email, userInfo.Name)

	// Creating Token if not exists
	var token models.Token
	tokenRes := database.DB.Where("token = ?", hashToken).First(&token)
	if tokenRes.Error != nil {
		token = models.Token{
			Token:  hashToken,
			UserId: &user.ID,
		}
		if err := database.DB.Create(&token).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
	var input models.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameters missing"})
		return
	}

	// Check if the user already exists
	var user models.User
	if err := database.DB.Where("username = ? OR email = ?", input.UsernameOrEmail, input.UsernameOrEmail).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while logging in"})
		}
		return
	}

	// Comparing Password Hash
	err := utils.CompareHash(user.Password, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generating Hash Token
	hashToken := utils.Hash(user.Email, input.Password)

	// Check if token exists
	var token models.Token
	tokenRes := database.DB.Where("token = ?", hashToken).First(&token)
	if tokenRes.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Signup(c *gin.Context) {
	var input models.SignupRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required parameters missing"})
		return
	}

	// Check if the user already exists
	var existingUser models.User
	if err := database.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Proceed with signup if the user does not already exist
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for existing user"})
			return
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{"error": "Username or email already in use"})
		return
	}

	// Hashing password
	hashedPassword := utils.Hash(input.Password)

	// Create the user in the database
	newUser := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user to the database"})
		return
	}

	// Generating Hash Token
	hashToken := utils.Hash(input.Email, input.Password)

	// Creating Token if not exists
	var token models.Token
	tokenRes := database.DB.Where("token = ?", hashToken).First(&token)
	if tokenRes.Error != nil {
		token = models.Token{
			Token:  hashToken,
			UserId: &newUser.ID,
		}
		if err := database.DB.Create(&token).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
