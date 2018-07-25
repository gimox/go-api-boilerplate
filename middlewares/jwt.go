package middlewares

import (
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"errors"
	"log"
	"net/http"
)

// middleware that add jwt Auth
func JwtAuth(secret string, decode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwtLib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 101, "error": "Authorization token failed"})
			c.AbortWithError(http.StatusUnauthorized, err)
			log.Println("token KO")
			return
		}

		c.Set("token", token) // add raw token, so controller can decode itself

		if decode {

			if decoded, ok := token.Claims.(jwtLib.MapClaims); ok && token.Valid {
				c.Set("jwt", decoded) // add decoded data for controller
				c.Next()

			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 102, "error": "Authorization token failed"})
				c.AbortWithError(http.StatusUnauthorized, errors.New("Authorization token failed!"))
				log.Println("Unable to decode Token")
			}

		}

	}
}
