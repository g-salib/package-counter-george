package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartApplication ... start the server and mapurl
func StartApplication() {
	log.Println("starting the server")
	server := gin.Default()
	server.Use(CORSMiddleware())

	// Serve static files (e.g., CSS and JavaScript) from the "static" directory.
	server.Static("/static", "./static")

	// Define a route to serve the HTML UI.
	server.LoadHTMLGlob("templates/*") // Load HTML templates from the "templates" directory.
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ui.html", nil)
	})

	mapUrls(server)
	server.Run(":5050")
}
