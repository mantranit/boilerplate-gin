package auth

import (
	"net/http"
	"time"

	"izihrm/models"
	"izihrm/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CtrlAuthenticate : get token with username/password
func CtrlAuthenticate(c *gin.Context) {
	var body FormLogin
	c.ShouldBind(&body)
	c.JSON(http.StatusOK, gin.H{
		"body": body,
	})

	expirationTime := time.Now().Add(23 * time.Hour)
	// Create the Claims
	claims := utils.CustomClaims{
		Role: utils.ADMIN,
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
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       token,
	})
}

// CtrlRegister ...
func CtrlRegister(c *gin.Context) {
	var rf FormRegister
	c.ShouldBind(&rf)

	hash, _ := bcrypt.GenerateFromPassword([]byte(rf.Password), bcrypt.DefaultCost)
	acc := models.Account{
		Username:  rf.Username,
		Email:     rf.Email,
		Hash:      string(hash),
		Status:    "PENDING",
		CreatedBy: rf.Username,
	}

	if obj := utils.DB.Create(&acc); obj.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "Success",
			"data":       acc,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    obj.Error.Error(),
		})
	}
}

// CtrlMe : get current account by token login
func CtrlMe(c *gin.Context) {
	claims := utils.GetClaims(c)
	accountID := claims.Issuer

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       accountID,
	})
}
