package infrastructure

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	domain "github.com/legend123213/go_togo/Task07/task-manager/Domain"
	"golang.org/x/crypto/bcrypt"
)

// Genratetoken generates a JWT token for the given user and password.
func Genratetoken(user *domain.User, pwd string) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

	// User login logic
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)) != nil {
		return "Invalid username or password", nil
	}

	if !CheckPassword(user, pwd) {
		return "Invalid username or password", nil
	}

	expirationTime := time.Now().Add(24 * 7 * time.Hour)
	claims := &domain.Claims{
		ID:       user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(jwtSecret)
	return jwtToken, err
}