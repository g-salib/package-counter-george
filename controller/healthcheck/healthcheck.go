package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping ...
func Ping(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"pong": "pong"},
	)
}
