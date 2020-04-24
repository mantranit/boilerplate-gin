package account

import (
	"izihrm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CtrlGetAll ...
func CtrlGetAll(c *gin.Context, db *gorm.DB) {
	var accounts []Account
	obj := db.Find(&accounts)
	if obj.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    obj.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       obj.Value,
	})
}

// CtrlGetByID ...
func CtrlGetByID(c *gin.Context, db *gorm.DB) {
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       accountID,
	})
}

// CtrlCreate ...
func CtrlCreate(c *gin.Context, db *gorm.DB) {
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       accountID,
	})
}

// CtrlUpdate ...
func CtrlUpdate(c *gin.Context, db *gorm.DB) {
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       accountID,
	})
}

// CtrlDelete ...
func CtrlDelete(c *gin.Context, db *gorm.DB) {
	accID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	var acc models.Account
	obj := db.Where("id = ?", accID).Find(&acc)
	if obj.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "NotFound: accountID",
		})
	} else {
		obj2 := db.Delete(obj.Value)
		if obj2.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusInternalServerError,
				"message":    obj2.Error.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusOK,
				"message":    "Success",
				"data":       obj2.Value,
			})
		}
	}
}
