package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Authenticator interface {
	GenerateToken(userId int64, email string) (string, error)
}

type authenticator struct {
	jwtSecret string
	jwtExp    time.Duration
}

func (a *authenticator) GenerateToken(userId int64, email string) (string, error) {
	if userId == 0 || email == "" {
		return "", errors.New("user id and email are required")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     jwt.NewNumericDate(time.Now().Add(a.jwtExp)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewAuthenticator(jwtSecret string, jwtExp time.Duration) Authenticator {
	return &authenticator{
		jwtSecret: jwtSecret,
		jwtExp:    jwtExp,
	}
}
