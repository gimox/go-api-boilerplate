package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/spf13/viper"

	"../middlewares"
	"../controllers"
)

func NewRouter(config *viper.Viper) *gin.Engine {
	router := gin.New()

	// middleware
	router.Use(cors.Default())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// base routes
	router.GET("/", controllers.Home)
	router.GET("/ping", controllers.Ping)
	router.GET("/health", controllers.Health)

	api := router.Group("/api")
	api.Use(middlewares.JwtAuth(config.GetString("jwt.secret"), true))
	{
		v1 := api.Group("/v1")
		v1.GET("/user", controllers.GetUser)
	}


	return router
}