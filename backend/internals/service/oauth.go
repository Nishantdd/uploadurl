package service

import (
	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/zalando/gin-oauth2/google"
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
