package account

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CtrlGetAll ...
func CtrlGetAll(c *gin.Context) {
	if obj := ServiceGetAll(c); obj.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "Success",
			"data":       obj.Value,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    obj.Error.Error(),
		})
	}
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
	if accID, err := strconv.ParseUint(c.Param("id"), 10, 64); err == nil {
		if obj := ServiceGetByID(uint(accID)); obj.RowsAffected > 0 {
			if obj2 := ServiceDelete(uint(accID)); obj2.Error == nil {
				c.JSON(http.StatusOK, gin.H{
					"statusCode": http.StatusOK,
					"message":    "Success",
					"data":       obj2.Value,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"statusCode": http.StatusInternalServerError,
					"message":    obj2.Error.Error(),
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusNotFound,
				"message":    "NotFound: accountID",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
	}
}
