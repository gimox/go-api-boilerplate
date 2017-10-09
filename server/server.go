package server

import (
	"github.com/gin-gonic/gin"
	"go-api-boilerplate/lib/config"
)


func Init() {
	configEnv := config.GetConfig() // get configuration enviroment ./config/yourenv

	// disable log per production mode
	if configEnv.GetBool("production") {
		gin.SetMode(gin.ReleaseMode)
	}


	r := NewRouter(configEnv) // start routes
	r.Run(":" + configEnv.GetString("server.port")) // start server
}