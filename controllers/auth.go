package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthSignIn(c *gin.Context) {

	var loginModel LoginCredentials

	if c.BindJSON(&loginModel) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 202, "message": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "hi", "credentials": loginModel})

	return
}
