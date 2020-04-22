package controllers

import (
	"net/http"

	"izihrm/utils"

	"github.com/gin-gonic/gin"
)

// UserController class
type UserController struct{}

// Current : get current user follow token login
func (user *UserController) Current(c *gin.Context) {
	claims := utils.GetClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       claims,
	})
}
