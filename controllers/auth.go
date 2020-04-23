package controllers

import (
	"net/http"
	"time"

	"izihrm/forms"
	"izihrm/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthController class
type AuthController struct{}

// Authenticate : get token with username/password
func (user *AuthController) Authenticate(c *gin.Context) {
	var body forms.Login
	c.ShouldBind(&body)
	c.JSON(http.StatusOK, gin.H{
		"body": body,
	})

	expirationTime := time.Now().Add(23 * time.Hour)
	// Create the Claims
	claims := forms.CustomClaims{
		Role: "ADMIN",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    body.Email,
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

// Me : get current user by token login
func (user *AuthController) Me(c *gin.Context) {
	claims := utils.GetClaims(c)
	userID := claims.Issuer

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       userID,
	})
}
