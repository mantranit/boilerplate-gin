package routers

import (
	"izihrm/controllers"
	"izihrm/utils"

	"github.com/gin-gonic/gin"
)

// APIUser ...
func APIUser(router *gin.Engine) {
	userCtrl := new(controllers.UserController)
	api := router.Group("/api")
	{
		apiUser := api.Group("/user")
		apiUser.Use(utils.Authorization("ADMIN"))
		apiUser.GET("/", userCtrl.GetUser)
		apiUser.POST("/", userCtrl.GetUser)
		apiUser.GET("/:id", userCtrl.GetUser)
		apiUser.POST("/:id", userCtrl.GetUser)
		apiUser.PUT("/:id", userCtrl.GetUser)
		apiUser.DELETE("/:id", userCtrl.GetUser)
	}
}
