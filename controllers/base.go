package controllers

import (
	"github.com/gin-gonic/gin"
)

func BaseHome(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hi"})
	return
}

func BasePing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
	return
}

func BaseHealth(c *gin.Context) {
	c.String(200, "1")
	return
}
