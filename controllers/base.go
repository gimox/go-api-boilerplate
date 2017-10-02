package controllers

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hi"})
	return
}


func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
	return
}
