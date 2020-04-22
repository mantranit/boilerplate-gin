package main

import (
	"net/http"
	"os"

	"izihrm/controllers"
	"izihrm/utils"

	"github.com/gin-gonic/gin"
)

// Note: This is just an example for a tutorial
// VALID_AUTHENTICATIONS is all role of user
var (
	VALID_AUTHENTICATIONS = []string{"user", "admin", "subscriber"}
)

func main() {

	router := gin.Default()

	authCtrl := new(controllers.AuthController)
	userCtrl := new(controllers.UserController)
	api := router.Group("/api")
	{
		api.POST("/authenticate", authCtrl.Authenticate)
	}
	{
		apiUser := api.Group("/user")
		apiUser.Use(utils.Authorization())
		{
			apiUser.GET("/current", userCtrl.Current)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "NotFound",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
