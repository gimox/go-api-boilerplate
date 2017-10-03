package middlewares

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"log"
	"errors"
)

func JwtAuth(secret string, decode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithError(401, err)
		}

		c.Set("token", token)

		if !decode {
			log.Println("Jwt NO decode", token)
			c.Next()
			return
		}

		log.Println("Jwt decode invoked", token)

		if decoded, ok := token.Claims.(jwt_lib.MapClaims); ok && token.Valid {
			c.Set("jwt", decoded)
			c.Next()
		} else {
			c.AbortWithError(401, errors.New("Authorization token failed!"))
		}

	}
}
