package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	//load all html files and components
	router.LoadHTMLFiles("templates/index.html", "templates/components/head.html")

	//load static directory
	router.StaticFS("static", http.Dir("./static"))

	router.SetTrustedProxies(nil)

	//site routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
