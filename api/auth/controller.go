package auth

import (
	"fmt"
	"net/http"
	"time"

	"izihrm/api/account"
	"izihrm/models"
	"izihrm/utils"
	"izihrm/utils/send"

	"github.com/cbroglie/mustache"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// CtrlAuthenticate : get token with username/password
func CtrlAuthenticate(c *gin.Context) {
	var login FormLogin
	c.ShouldBind(&login)
	err := utils.Validate.Struct(login)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	var account models.Account
	result := utils.DB.Where("email like ?", login.Email).Find(&account)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "NotFound",
		})
		return
	}

	if account.Status != models.StatusActive {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusLocked,
			"message":    "Locked: Account not active",
		})
		return
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
		AccountID: account.ID,
		Role:      utils.ADMIN,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    account.Email,
			Subject:   account.Hash,
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
	err := utils.Validate.Struct(register)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
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
		Email:     register.Email,
		Hash:      string(hash),
		Status:    models.StatusActive,
		CreatedBy: "register",
	}

	result := utils.DB.Create(&account)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
	})
}

// CtrlForgotPassword ...
func CtrlForgotPassword(c *gin.Context) {
	var forgotPassword FormForgotPassword
	c.ShouldBind(&forgotPassword)
	err := utils.Validate.Struct(forgotPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	var account models.Account
	obj := utils.DB.Where("email like ?", forgotPassword.Email).Find(&account)
	if obj.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "NotFound",
		})
		return
	}

	token := password.MustGenerate(64, 10, 0, false, true)

	account.Token = token
	result := utils.DB.Save(&account)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    result.Error.Error(),
		})
		return
	}

	var url = fmt.Sprintf("%s/password/%s", utils.ViperEnvVariable("DOMAIN_FE"), token)
	content, errContent := mustache.RenderFile("templates/emails/RecoverPassword.mustache", map[string]string{"url": url})
	if errContent != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    errContent.Error(),
		})
		return
	}

	responseBody := send.Mail(content, forgotPassword.Email, "Reset your password")
	if responseBody != "" {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    responseBody,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
	})
}

// CtrlResetPassword ...
func CtrlResetPassword(c *gin.Context) {
	token := c.Param("token")
	var resetPassword FormResetPassword
	c.ShouldBind(&resetPassword)
	err := utils.Validate.Struct(resetPassword)
	if err != nil || token == "" {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	var account models.Account
	obj := utils.DB.Where("token like ?", token).Find(&account)
	if obj.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "NotFound",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(resetPassword.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    err.Error(),
		})
		return
	}

	account.Hash = string(hash)
	result := utils.DB.Save(&account)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
	})
}

// CtrlMe : get current account by token login
func CtrlMe(c *gin.Context) {
	claims := utils.GetClaims(c)

	var data account.Account
	result := utils.DB.Where("id = ?", claims.AccountID).First(&data)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNoContent,
			"message":    result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       data,
	})
}
