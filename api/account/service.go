package account

import (
	"izihrm/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ServiceGetAll ...
func ServiceGetAll(c *gin.Context) *gorm.DB {
	var accounts []Account
	result := utils.DB.Find(&accounts)
	return result
}

// ServiceGetByID ...
func ServiceGetByID(ID uint) *gorm.DB {
	var account Account
	result := utils.DB.Where("id = ?", ID).Find(&account)
	return result
}

// ServiceDelete ...
func ServiceDelete(ID uint) *gorm.DB {
	account := &Account{
		Model: gorm.Model{
			ID: ID,
		},
	}
	result := utils.DB.Delete(&account)
	return result
}
