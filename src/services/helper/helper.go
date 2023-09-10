package helper

import (
	"net/http"
	"strconv"
	helperResponse "app-aposrgb/src/helpers/response"

	"github.com/gin-gonic/gin"
)

func GetIntIdByParams(c *gin.Context) (int, bool){
	stringId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, helperResponse.NewBadRequestResponse("Bad request, not found require params"))
		return 0, false
	}
	id, err  := strconv.Atoi(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperResponse.NewBadRequestResponse("Invalid id params"))
		return 0, false
	}
	return id, true
}