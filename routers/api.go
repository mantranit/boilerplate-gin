package routers

import (
	"izihrm/controllers"
	"izihrm/utils"

	"github.com/gin-gonic/gin"
)

// API ...
func API(router *gin.Engine) {
	authCtrl := new(controllers.AuthController)
	api := router.Group("/api")
	{
		api.POST("/authenticate", authCtrl.Authenticate)
	}
	{
		api.Use(utils.Authorization())
		api.GET("/me", authCtrl.Me)
	}
}
