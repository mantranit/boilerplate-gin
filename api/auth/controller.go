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
	var login FormLogin
	c.ShouldBind(&login)

	var account models.Account
	obj := utils.DB.Where("username like ?", login.Email).Find(&account)
	if obj.RowsAffected == 0 {
		obj = utils.DB.Where("email like ?", login.Email).Find(&account)
		if obj.RowsAffected == 0 {
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusNotFound,
				"message":    "NotFound",
			})
			return
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Hash), []byte(login.Password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNotAcceptable,
			"message":    "NotAcceptable: Wrong password",
		})
		return
	}

	expirationTime := time.Now().Add(23 * time.Hour)
	// Create the Claims
	claims := utils.CustomClaims{
		Role: utils.ADMIN,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    login.Email,
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
	} else {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "Success",
			"data":       token,
		})
	}
}

// CtrlRegister ...
func CtrlRegister(c *gin.Context) {
	var register FormRegister
	c.ShouldBind(&register)

	if register.Email == "" || register.Password == "" || !utils.ValidateEmail(register.Email) {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "BadRequest",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    err.Error(),
		})
		return
	}
	account := models.Account{
		Username:  register.Username,
		Email:     register.Email,
		Hash:      string(hash),
		Status:    "PENDING",
		CreatedBy: register.Username,
	}

	if obj := utils.DB.Create(&account); obj.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    obj.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "Success",
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
