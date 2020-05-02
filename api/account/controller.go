package account

import (
	"fmt"
	"izihrm/models"
	"izihrm/utils"
	"izihrm/utils/send"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// CtrlGetAll ...
func CtrlGetAll(c *gin.Context) {
	var accounts []Account
	result := utils.DB.Find(&accounts)

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
		"data":       accounts,
	})
}

// CtrlGetByID ...
func CtrlGetByID(c *gin.Context) {
	accID := c.Param("id")
	var account Account
	result := utils.DB.Where("id = ?", accID).Find(&account)

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
		"data":       account,
	})
}

// CtrlCreate ...
func CtrlCreate(c *gin.Context) {
	var account models.Account
	c.ShouldBind(&account)

	err := utils.Validate.Struct(account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	pwd := password.MustGenerate(16, 4, 4, false, false)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    err.Error(),
		})
		return
	}

	token := password.MustGenerate(64, 10, 0, false, true)

	account.Status = models.StatusPending
	account.Hash = string(hash)
	account.CreatedBy = utils.GetClaims(c).Issuer
	account.Token = token

	result := utils.DB.Create(&account)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    result.Error.Error(),
		})
		return
	}

	var url = fmt.Sprintf("%s/password/%s", utils.ViperEnvVariable("DOMAIN_FE"), token)
	content, errContent := mustache.RenderFile("templates/emails/ActivateAccount.mustache", map[string]string{"url": url})
	if errContent != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    errContent.Error(),
		})
		return
	}

	responseBody := send.Mail(content, account.Email, "Activate your account")
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

// CtrlUpdate ...
func CtrlUpdate(c *gin.Context) {
	accID := c.Param("id")
	var account Account
	result := utils.DB.Where("id = ?", accID).Find(&account)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    result.Error.Error(),
		})
		return
	}

	c.ShouldBind(&account)
	account.UpdatedBy = utils.GetClaims(c).Issuer
	result = utils.DB.Save(&account)
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
		"data":       account,
	})
}

// CtrlDelete ...
func CtrlDelete(c *gin.Context) {
	accID := c.Param("id")
	var account Account
	result := utils.DB.Where("id = ?", accID).Find(&account)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    result.Error.Error(),
		})
		return
	}
	result = utils.DB.Unscoped().Delete(&account)

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
