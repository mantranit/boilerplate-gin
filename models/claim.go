package models

import "github.com/dgrijalva/jwt-go"

// CustomClaims for token
type CustomClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
