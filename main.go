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
	db := utils.ConnectDatabase()
	defer db.Close()
	db.AutoMigrate(&models.Account{})

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(utils.ErrorHandler)

	// set CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// defined all routes
	auth.SetupRouter(r, db)
	account.SetupRouter(r, db)

	// fallback route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "NotFound",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
