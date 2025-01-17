package auth

import (
	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func NewAuth() {
	cfg := config.Load()

	key := cfg.OAuth.SessionSecret
	MaxAge := 86400 * 30
	IsProd := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store
	goth.UseProviders(
		google.New(cfg.OAuth.GoogleClientId, cfg.OAuth.GoogleClientSecret, "http://localhost:3000/api/auth/google/callback"),
	)

}
