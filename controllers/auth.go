package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthSignIn(c *gin.Context) {

	var loginCrd LoginCredentials

	if c.BindJSON(&loginCrd) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 202, "message": "invalid username or password"})
		return
	}

	if loginCrd.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 202, "message": "Username can not be blank"})
		return
	}

	if loginCrd.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 202, "message": "Password can not be blank"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "hi", "credentials": loginCrd})

	return
}
