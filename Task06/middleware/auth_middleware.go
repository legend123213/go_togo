package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/legend123213/go_togo/Task06/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Claims struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `json:"username`
	IsAdmin bool    `json:"role"`
	jwt.StandardClaims
}

func AuthMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {

	var jwtSecret = []byte("your_jwt_secret")
	authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
	tokenString := strings.Split(authHeader, " ")[1]
    // Parse the token
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        c.JSON(http.StatusForbidden,err)
    }

    // Extract the claims
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		id := claims.ID
		
		user := &models.User{
			ID:       id,
			Username: claims.Username,
			IsAdmin:  claims.IsAdmin,
		}
		c.Set("id",user.ID)
		c.Set("username",user.Username)
		c.Set("isActive",user.IsAdmin)
      c.Next()
    }

    c.JSON(http.StatusForbidden,errors.New("invalid token")) 
	 c.Abort()
	 return 
	}

	
}

func AdminMiddleware() gin.HandlerFunc{
	return func (c *gin.Context)  {
		isadmin, _ := c.MustGet("isActive").(bool)
		log.Println(isadmin)
		if isadmin {
			c.Next()
			
		}else{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"it is admin role"})
		c.Abort()
		return 
		}
		
		
	}
}