package service

import (
	"errors"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/Nishantdd/uploadurl/backend/internals/database"
	"github.com/Nishantdd/uploadurl/backend/internals/models"
)

func RegisterUrl(originalUrl string, slugValue string, urlType string, userId *uint64) error {

	shortUrl := config.Load().Server.DomainAddress + "/" + slugValue

	url := models.Url{
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
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
