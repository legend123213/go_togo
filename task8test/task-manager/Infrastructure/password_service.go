package infrastructure

import (
	domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
	"golang.org/x/crypto/bcrypt"
)

// CheckPassword checks if the provided password matches the user's hashed password.
func CheckPassword(user *domain.User, pwd string) bool {
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)) != nil {
		return false
	}
	return true
}