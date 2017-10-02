package server

import "ginapi/lib/config"

func Init() {
	config := config.GetConfig()

	r := NewRouter(config)
	r.Run(":" + config.GetString("server.port"))
}