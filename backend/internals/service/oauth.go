package service

import (
	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Oauth2Config oauth2.Config
var Oauth2State string

func InitOauth() (oauth2.Config, string) {
	cfg := config.Load()
	Oauth2Config = oauth2.Config{
		ClientID:     cfg.OAuth.GoogleClientID,
		ClientSecret: cfg.OAuth.GoogleClientSecret,
		RedirectURL:  cfg.OAuth.RedirectURL,
		Scopes:       []string{"openid", "profile", "email"}, // Permissions requested
		Endpoint:     google.Endpoint,
	}

	Oauth2State = utils.GenerateState()
	return Oauth2Config, Oauth2State
}
