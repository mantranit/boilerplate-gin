package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController class
type UserController struct{}

// GetUser : get user by userId
func (user *UserController) GetUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       userID,
	})
}

// Delete : userId
func (user *UserController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"params":     c.Params,
	})
}
