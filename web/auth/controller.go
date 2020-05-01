package auth

import (
	"fmt"
	"izihrm/api/auth"
	"izihrm/models"
	"izihrm/utils"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var error string

// CtrlResetPassword ...
func CtrlResetPassword(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "resetPassword", gin.H{
		"title": "Reset your password",
		"token": c.Param("token"),
		"error": error,
	})
}

// CtrlCreatePassword ...
func CtrlCreatePassword(c *gin.Context) {
	error = "" // reset error
	var token = c.Param("token")
	var resetPassword = &auth.FormResetPassword{
		Password:        c.PostForm("password"),
		ConfirmPassword: c.PostForm("confirmPassword"),
	}
	err := utils.Validate.Struct(resetPassword)
	if err != nil {
		error = err.Error()
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/reset-password/%s", token))
		return
	}

	var account models.Account
	obj := utils.DB.Where("token like ?", token).Find(&account)
	if obj.RowsAffected == 0 {
		error = "NotFound"
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/reset-password/%s", token))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(resetPassword.Password), bcrypt.DefaultCost)
	if err != nil {
		error = err.Error()
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/reset-password/%s", token))
		return
	}

	account.Hash = string(hash)
	account.Token = ""
	result := utils.DB.Save(&account)
	if result.Error != nil {
		error = result.Error.Error()
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/reset-password/%s", token))
		return
	}

	error = ""
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/success"))
}

// CtrlSuccess ...
func CtrlSuccess(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "success", gin.H{
		"title": "Success",
	})
}