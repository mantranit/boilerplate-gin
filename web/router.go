package web

import (
	"html/template"
	"izihrm/web/auth"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine) {
	// static files
	router.Static("/assets", "./assets")

	// new template engine
	router.HTMLRender = ginview.New(goview.Config{
		Root:      "templates/views",
		Extension: ".html",
		Master:    "layouts/master",
		// Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	web := router.Group("/")
	{
		web.GET("/password/:token", auth.CtrlPassword)
		web.POST("/create-password/:token", auth.CtrlCreatePassword)
		web.GET("/success", auth.CtrlSuccess)
	}
}
