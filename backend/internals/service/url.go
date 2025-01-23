package service

import (
	"errors"
	"math/rand/v2"
	"strings"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
)

// Takes length as input and returns a alphanumeric string of given length
func GenerateRandomSlug(len int) string {
	var slug strings.Builder

	allowedCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1233456789"
	for i := 0; i < len; i++ {
		slug.WriteByte(allowedCharacters[rand.IntN(16)])
	}

	return slug.String()
}

func RegisterUrl(originalUrl string, slugValue string, urlType string, userId *uint64) error {
	shortUrl := config.Load().Server.DomainAddress + "/" + slugValue

	url := models.Url{
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
		Type:        urlType,
		Slug:        slugValue,
		UserId:      userId,
	}

	res := database.DB.Create(&url)
	if res.Error != nil {
		return errors.New("failed to create URL: " + res.Error.Error())
	}

	slug := models.Slug{
		UrlId:  url.ID,
		Slug:   slugValue,
		UserId: userId,
	}

	res = database.DB.Create(&slug)
	if res.Error != nil {
		return errors.New("failed to create slug: " + res.Error.Error())
	}

	return nil
}
