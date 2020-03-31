package payor

import (
	"errors"
	"net/http"
	"os"
	"strings"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// BearerTokenAuthMiddleware : middleware
func BearerTokenAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the token from the header
		if c.FullPath() == "/" {
			c.Next()
		} else if c.FullPath() == "/auth/login" {
			c.Set("DB", db)
			c.Next()
		} else {
			authHeader := c.GetHeader("Authorization")
			bearer := strings.Replace(authHeader, "Bearer ", "", -1)

			token, err := jwt_lib.Parse(string(bearer), func(token *jwt_lib.Token) (interface{}, error) {
				if token.Header["alg"] != "HS256" {
					return nil, errors.New("Unexpected signing Method")
				}
				b := ([]byte(os.Getenv("JWT_SECRET")))
				return b, nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
				c.Abort()
				return
			}
			if token.Valid {
				c.Set("DB", db)
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
				c.Abort()
				return
			}

		}
	}
}
