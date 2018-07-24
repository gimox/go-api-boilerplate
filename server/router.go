package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/spf13/viper"
	"go-api-boilerplate/middlewares"
	"go-api-boilerplate/controllers"
)

func NewRouter(config *viper.Viper) *gin.Engine {
	router := gin.New()

	// Global middleware
	router.Use(cors.Default())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// base routes
	router.GET("/", controllers.BaseHome)
	router.GET("/ping", controllers.BasePing)
	router.GET("/health", controllers.BaseHealth)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	// Authorization routes
	auth := v1.Group("/auth")
	{
		auth.POST("/signIn", controllers.AuthSignIn)
	}

	// route protected with JWT
	private := v1.Group("")
	private.Use(middlewares.JwtAuth(config.GetString("jwt.secret"), true))
	{
		private.GET("/user", controllers.GetUser)
	}

	return router
}
