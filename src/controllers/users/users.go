package users

import (
	helperResponse "app-aposrgb/src/helpers/response"
	"app-aposrgb/src/repository"
	"app-aposrgb/src/services/helper"
	"app-aposrgb/src/services/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	response.SetHeadersJson(c)
	c.JSON(http.StatusOK, repository.GetUsers())
}

func GetUser(c *gin.Context) {
	response.SetHeadersJson(c)
	id, ok := helper.GetIntIdByParams(c)
	if !ok {
		return
	}
	user, ok := repository.GetUser(id)
	if !ok {
		c.JSON(http.StatusNotFound, helperResponse.NewNotFoundResponse("User not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}