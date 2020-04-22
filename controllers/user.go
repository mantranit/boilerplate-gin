package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mantranit/izihrm/models"
	"github.com/mantranit/izihrm/utils"
)

// Authenticate is function
func Authenticate(c *gin.Context) {
	expirationTime := time.Now().Add(23 * time.Hour)
	// Create the Claims
	claims := models.CustomClaims{
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "test",
		},
	}
	mySigningKey := []byte(utils.ViperEnvVariable("JWT_SECRET_KEY"))
	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signature.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Internal Server Error",
			"data":       token,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       token,
	})
}

// Current user
func Current(c *gin.Context) {
	claims := utils.GetClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       claims,
	})
}
