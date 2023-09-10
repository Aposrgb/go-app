package routes

import (
	"app-aposrgb/src/controllers/users"
	helperResponse "app-aposrgb/src/helpers/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	SetUserRoutes(router)
	SetAnyRoutes(router)
}

func SetUserRoutes(router *gin.Engine) {
	router.GET("/user", users.GetUsers)
	router.GET("/user/:id", users.GetUser)
}

func SetAnyRoutes(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helperResponse.NewNotFoundResponse("Not found method"))
	})
}
