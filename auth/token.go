package auth

import (
	"time"

	"github.com/azkazkazka/task-todo/config"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID string) (string, error) {
	conf := config.GetConfig()
	var jwtKey = []byte(conf.JWT_TOKEN_KEY)

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
