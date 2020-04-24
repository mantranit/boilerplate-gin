package account

import (
	"izihrm/utils"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		apiAccount := api.Group("/account")
		apiAccount.Use(utils.Authorization(utils.ADMIN))
		apiAccount.GET("/", CtrlGetAll)
		apiAccount.POST("/", CtrlCreate)
		apiAccount.GET("/:id", CtrlGetByID)
		apiAccount.PUT("/:id", CtrlUpdate)
		apiAccount.DELETE("/:id", CtrlDelete)
	}
}
