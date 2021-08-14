package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	errInvalidSignKey = errors.New("invalid token sign key")
	errInvalidToken   = errors.New("session token is invalid or expired")
)

// GenerateToken return a signed JWT token with the user permissions
func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.APISecret)
}

func ValidateToken(r *http.Request) error {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return errInvalidToken
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errInvalidToken
	}

	return nil
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) != 2 {
		return ""
	}

	return strings.Split(token, " ")[1]
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errInvalidSignKey
	}

	return config.APISecret, nil
}

// GetUserID returns the user ID stored in the token
func GetUserID(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return 0, errInvalidToken
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errInvalidToken
	}

	userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
	if err != nil {
		return 0, errInvalidToken
	}

	return userID, nil
}
