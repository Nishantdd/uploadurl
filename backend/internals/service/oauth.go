package service

import (
	"encoding/json"
	"net/http"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Oauth2Config oauth2.Config
var Oauth2State string

func InitOauth() (oauth2.Config, string) {
	cfg := config.Load()
	config := oauth2.Config{
		ClientID:     cfg.OAuth.GoogleClientID,
		ClientSecret: cfg.OAuth.GoogleClientSecret,
		RedirectURL:  cfg.OAuth.RedirectURL,
		Scopes:       []string{"openid", "profile", "email"}, // Permissions requested
		Endpoint:     google.Endpoint,
	}

	state := utils.GenerateState()
	return config, state
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
