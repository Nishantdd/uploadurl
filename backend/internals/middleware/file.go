package middleware

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func StoreMultipartFilesLocally() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			c.Abort()
		}

		filename := filepath.Base(file.Filename)
		savepath := "../../tmp/" + filename
		err = c.SaveUploadedFile(file, savepath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
		}

		c.Set("path", savepath)
		c.Next()
	}
}
