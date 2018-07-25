// implements routes
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-api-boilerplate/middlewares"
	"go-api-boilerplate/controllers"
	"github.com/gin-contrib/cors"
	_ "go-api-boilerplate/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

)

// @title API Swagger
// @version 1.0
// @description This is api sample doc.
// @termsOfService http://swagger.io/terms/
// @contact.name Giorgio Modoni
// @contact.url http://
// @contact.email modogio@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func NewRouter(config *viper.Viper) *gin.Engine {
	router := gin.New()

	// Global middleware
	router.Use(cors.Default())

	// router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// base routes
	router.GET("/", controllers.BaseHome)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/ping", controllers.BasePing)
	v1.GET("/health", controllers.BaseHealth)

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
