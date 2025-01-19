package service

import (
	"encoding/json"
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"golang.org/x/oauth2/google"
)

func InitOAuth() {
	cfg := config.Load()
	var (
		redirectURL  = cfg.OAuth.RedirectURL
		clientId     = cfg.OAuth.GoogleClientId
		clientSecret = cfg.OAuth.GoogleClientSecret
	)

	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
	secret := []byte(cfg.OAuth.OAuthSecret)

	// Initialize Google OAuth2
	google.SetupFromString(redirectURL, clientId, clientSecret, scopes, secret)
}

func GetGoogleUserInfo(client *http.Client) (*models.GoogleUserInfo, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo models.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
