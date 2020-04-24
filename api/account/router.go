package account

import (
	"izihrm/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		apiAccount := api.Group("/account")
		apiAccount.Use(utils.Authorization(utils.ADMIN))
		apiAccount.GET("/", func(c *gin.Context) {
			CtrlGetAll(c, db)
		})
		apiAccount.POST("/", func(c *gin.Context) {
			CtrlCreate(c, db)
		})
		apiAccount.GET("/:id", func(c *gin.Context) {
			CtrlGetByID(c, db)
		})
		apiAccount.PUT("/:id", func(c *gin.Context) {
			CtrlUpdate(c, db)
		})
		apiAccount.DELETE("/:id", func(c *gin.Context) {
			CtrlDelete(c, db)
		})
	}
}
