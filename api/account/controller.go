package account

import (
	"izihrm/utils"
	"net/http"
	"strconv"

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
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       accountID,
	})
}

// CtrlCreate ...
func CtrlCreate(c *gin.Context) {
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       accountID,
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
	accID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

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
