package controllers

import (
	"github.com/gin-gonic/gin"
)

// home route
// @Summary Show home route
// @Description this is the home of API
// @Accept  json
// @Produce  json
// @Router / [get]
func BaseHome(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hi"})
	return
}

// Ping
// @Summary Ping response
// @Description this is the home of API
// @Accept  json
// @Produce  json
// @Router /v1/ping [get]
func BasePing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
	return
}

// health response
// @Summary health response
// @Description like a ping with a simple response
// @Produce  json
// @Success 200 {string} string	"1"
// @Router /v1/health [get]
func BaseHealth(c *gin.Context) {
	c.String(200, "1")
	return
}
