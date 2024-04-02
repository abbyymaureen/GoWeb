package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Load HTML templates
	router.LoadHTMLGlob("templates/*.html")

	// Define route for root URL (index.html page)
	router.GET("/", func(c *gin.Context) {
		// Render index.html template
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// Define routes for other HTML pages
	router.GET("/lessTravelled.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lessTravelled.html", gin.H{
			"title": "Less Travelled",
		})
	})

	router.GET("/moreTravelled.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "moreTravelled.html", gin.H{
			"title": "More Travelled",
		})
	})

	router.GET("/talkWhale.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "talkWhale.html", gin.H{
			"title": "Talk to Whale",
		})
	})

	router.GET("/runFromWhale.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "runFromWhale.html", gin.H{
			"title": "Run from Whale",
		})
	})

	router.GET("/talkStranger.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "talkStranger.html", gin.H{
			"title": "Talk to Stranger",
		})
	})

	router.GET("/runStranger.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "runStranger.html", gin.H{
			"title": "Run from Stranger",
		})
	})

	router.GET("/runTiger.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "runTiger.html", gin.H{
			"title": "Run from Tiger",
		})
	})

	router.GET("/talkTiger.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "talkTiger.html", gin.H{
			"title": "Talk to Tiger",
		})
	})

	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main Page",
		})
	})

	// Start the server
	router.Run(":8080")
}
