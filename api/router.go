package api

import (
	"izihrm/api/account"
	"izihrm/api/auth"
	"izihrm/utils"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/authenticate", auth.CtrlAuthenticate)
		api.POST("/register", auth.CtrlRegister)
		api.POST("/forgot-password", auth.CtrlForgotPassword)
		api.POST("/reset-password", auth.CtrlResetPassword)
	}
	{
		api.Use(utils.Authorization())
		api.GET("/me", auth.CtrlMe)
	}
	{
		apiAccount := api.Group("/account")
		apiAccount.Use(utils.Authorization(utils.ADMIN))
		apiAccount.GET("/", account.CtrlGetAll)
		apiAccount.POST("/", account.CtrlCreate)
		apiAccount.GET("/:id", account.CtrlGetByID)
		apiAccount.PUT("/:id", account.CtrlUpdate)
		apiAccount.DELETE("/:id", account.CtrlDelete)
	}
}
