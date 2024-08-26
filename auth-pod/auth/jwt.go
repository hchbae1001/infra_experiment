package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var accessSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
var refreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID int) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(accessSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userID int) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAccessToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func ValidateRefreshToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
