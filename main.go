package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	// provided _any_ route, send the user to pywebdev.com
	// https://github.com/gin-gonic/gin/issues/1207
	router.Any("/*path", func(c *gin.Context) {
		// StatusSeeOther == HTTP 303
		// https://golang.org/src/net/http/status.go
		// https://github.com/gin-gonic/gin#redirects
		c.Redirect(http.StatusSeeOther, "https://www.pywebdev.com/")
	})

	router.Run(":" + port)
}
