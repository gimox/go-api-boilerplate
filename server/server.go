// start server
package server

import (
	"github.com/gin-gonic/gin"
	"go-api-boilerplate/lib/config"
)

// main function for server start
func Init() {
	configEnv := config.GetConfig() // get configuration enviroment ./config/yourenv

	// disable log per production mode
	if configEnv.GetBool("production") {
		gin.SetMode(gin.ReleaseMode)
	}

	// init router
	r := NewRouter(configEnv)                       // start routes
	r.Run(":" + configEnv.GetString("server.port")) // start server
}
