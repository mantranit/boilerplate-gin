package utils

import (
	"net/http"
	"strings"

	"izihrm/forms"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetClaims from token
func GetClaims(c *gin.Context) *forms.CustomClaims {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	reqToken = strings.TrimSpace(splitToken[1])

	token, _ := jwt.ParseWithClaims(reqToken, &forms.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(ViperEnvVariable("JWT_SECRET_KEY")), nil
	})

	claims, _ := token.Claims.(*forms.CustomClaims)

	return claims
}

// Authorization for new
func Authorization(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		if reqToken == "" || !strings.Contains(reqToken, "Bearer") || len(strings.Split(reqToken, ".")) != 3 {
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusBadRequest,
				"message":    "BadRequest",
			})
			c.Abort()
			return
		}
		splitToken := strings.Split(reqToken, "Bearer")
		reqToken = strings.TrimSpace(splitToken[1])

		token, err := jwt.ParseWithClaims(reqToken, &forms.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(ViperEnvVariable("JWT_SECRET_KEY")), nil
		})

		if !token.Valid {
			ve, _ := err.(*jwt.ValidationError)
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.JSON(http.StatusOK, gin.H{
					"statusCode": http.StatusNotAcceptable,
					"message":    "NotAcceptable: That's not even a token",
				})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.JSON(http.StatusOK, gin.H{
					"statusCode": http.StatusNotAcceptable,
					"message":    "NotAcceptable: Token is either expired or not active yet",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"statusCode": http.StatusNotAcceptable,
					"message":    "NotAcceptable: Couldn't handle this token",
				})
			}
			c.Abort()
			return
		}
		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		c.Next()
	}
}
