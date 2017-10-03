package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := c.MustGet("jwt")

	c.JSON(200, gin.H{"message": "user getUser", "user": user})
	return
}
