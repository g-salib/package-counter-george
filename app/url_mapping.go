package app

import (
	"github.com/gin-gonic/gin"
	"github.com/package-counter/controller/calculatepacks"
	"github.com/package-counter/controller/healthcheck"
)

func mapUrls(server *gin.Engine) {
	// health check
	healthGroup := server.Group("/api/v1")
	healthGroup.GET("/ping", healthcheck.Ping)
	packageValidator := server.Group("/api/v1/packs")
	packageValidator.GET("/calculate", calculatepacks.GET)

}
