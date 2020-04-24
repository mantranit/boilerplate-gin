package main

import (
	"net/http"
	"os"

	"izihrm/api/account"
	"izihrm/api/auth"
	"izihrm/models"
	"izihrm/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.DB = utils.ConnectDatabase()
	defer utils.DB.Close()
	utils.DB.AutoMigrate(&models.Account{})

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(utils.ErrorHandler)

	// set CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// defined all routes
	auth.SetupRouter(router)
	account.SetupRouter(router)

	// fallback route
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
