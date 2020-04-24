package auth

import (
	"izihrm/utils"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/authenticate", CtrlAuthenticate)
		api.POST("/register", CtrlRegister)
	}
	{
		api.Use(utils.Authorization())
		api.GET("/me", CtrlMe)
	}
}
