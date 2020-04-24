package auth

import (
	"izihrm/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		api.POST("/authenticate", func(c *gin.Context) {
			CtrlAuthenticate(c, db)
		})
		api.POST("/register", func(c *gin.Context) {
			CtrlRegister(c, db)
		})
	}
	{
		api.Use(utils.Authorization())
		api.GET("/me", func(c *gin.Context) {
			CtrlMe(c, db)
		})
	}
}
