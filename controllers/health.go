package controllers

import (
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(200, "1")
	return
}
