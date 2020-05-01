package web

import (
	"html/template"
	"izihrm/web/auth"
	"time"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine) {
	// static files
	router.Static("/assets", "./assets")

	// new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "templates/views",
		Extension: ".html",
		Master:    "layouts/master",
		// Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	web := router.Group("/")
	{
		web.GET("/reset-password/:token", auth.CtrlResetPassword)
		web.POST("/create-password/:token", auth.CtrlCreatePassword)
		web.GET("/success", auth.CtrlSuccess)
	}
}
