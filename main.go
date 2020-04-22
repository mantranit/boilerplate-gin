package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mantranit/izihrm/controllers"
	"github.com/mantranit/izihrm/utils"
)

// Note: This is just an example for a tutorial
// VALID_AUTHENTICATIONS is all role of user
var (
	VALID_AUTHENTICATIONS = []string{"user", "admin", "subscriber"}
)

func main() {

	router := gin.Default()

	api := router.Group("/api")
	// no authentication endpoints
	{
		api.POST("/authenticate", controllers.Authenticate)
	}
	{
		apiUser := api.Group("/user")
		apiUser.Use(utils.Authorization())
		{
			apiUser.GET("/current", controllers.Current)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
