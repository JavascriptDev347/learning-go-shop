// Package utils for handling jwt token related operations
package utils

import (
	"errors"
	"time"

	"github.com/JavascriptDev347/learning-go-shop/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// Claims struct for token
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateTokenPair for createing refresh and accessToken
func GenerateTokenPair(cfg *config.JWTConfig, userID uint, email string, role string) (accessToken, refreshToken string, err error) {
	// create an access token
	accessClaims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.ExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.ExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// ValidateToken for validating the token
func ValidateToken(tokenString, secret string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
