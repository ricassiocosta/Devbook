package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken return a signed JWT token with the user permissions
func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("not-so-secret"))
}
