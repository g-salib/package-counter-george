package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

// StartApplication ... start the server and mapurl
func StartApplication() {
	// mapUrls()
	log.Println("startig the server")
	server := gin.Default()
	server.Use(CORSMiddleware())
	mapUrls(server)
	server.Run(":5050")
}
