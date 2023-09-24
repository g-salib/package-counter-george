package calculatepacks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/package-counter/domain/calculation"
	"github.com/package-counter/utils/httpresponse"
)

func GET(context *gin.Context) {

	orderQuantity := context.Query("quantity")
	qty, _ := strconv.Atoi(orderQuantity)
	packs := calculation.CalculatePacks(float64(qty))

	// response
	response := httpresponse.NewBaseResponseWithoutData(packs, "Pack Fetched Successfully", "GET")

	context.JSON(
		http.StatusOK,
		response,
	)

}
