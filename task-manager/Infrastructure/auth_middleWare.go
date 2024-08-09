package infrastructure

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	domain "github.com/legend123213/go_togo/Task07/task-manager/Domain"
)

// AuthMiddleware is a middleware function that handles authentication.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusForbidden, err)
			return
		}
		if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
			// Check expiration
			if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
				c.Abort()
				return
			}

			// Extract the claims
			if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
				id := claims.ID

				user := &domain.User{
					ID:       id,
					Username: claims.Username,
					IsAdmin:  claims.IsAdmin,
				}
				c.Set("id", user.ID)
				c.Set("username", user.Username)
				c.Set("isActive", user.IsAdmin)
				c.Next()
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
	}
}

// AdminMiddleware is a middleware function that checks if the user is an admin.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isadmin := c.MustGet("isActive").(bool)
		if isadmin {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "it is admin role"})
		}
	}
}
