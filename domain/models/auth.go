package models

import "github.com/golang-jwt/jwt"

type Auth struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
	UserID      uint   `json:"user_id"`
}

type JWTClaims struct {
	jwt.StandardClaims
}
