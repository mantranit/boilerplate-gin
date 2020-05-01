package account

import (
	"izihrm/models"
	"izihrm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       account,
	})
}

// CtrlUpdate ...
func CtrlUpdate(c *gin.Context) {
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       accountID,
	})
}

// CtrlDelete ...
func CtrlDelete(c *gin.Context) {
	accID := c.Param("id")
	dbResult := utils.DB.Where("id = ?", accID).Delete(&Account{})

	if dbResult.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNoContent,
			"message":    "NoContent",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
	})
}
