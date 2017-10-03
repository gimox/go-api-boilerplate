package controllers

import "github.com/gin-gonic/gin"

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func AuthSignIn(c *gin.Context) {

	var loginCrd LoginCredentials
	c.BindJSON(&loginCrd)

	c.JSON(200, gin.H{"message": "hi", "credentials": loginCrd})
	return
}
