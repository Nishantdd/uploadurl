package service

import (
	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/zalando/gin-oauth2/google"
)

func InitOAuth() {
	cfg := config.Load()

	redirectURL := cfg.OAuth.RedirectURL
	credFile := cfg.OAuth.CredFilePath
	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
	secret := []byte(cfg.OAuth.OAuthSecret)

	// Initialize Google OAuth2
	google.Setup(redirectURL, credFile, scopes, secret)
}
