package calculatepacks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/package-counter/domain/calculation"
	"github.com/package-counter/utils/httpresponse"
)

func GET(context *gin.Context) {

	orderQuantity := context.Query("quantity")
	packs := calculation.CalculatePacks(orderQuantity)
	// response
	response := httpresponse.NewBaseResponseWithoutData(packs, "Pack Fetched Successfully", "POST")

	context.JSON(
		http.StatusOK,
		response,
	)
}
