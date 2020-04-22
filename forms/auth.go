package forms

import "github.com/dgrijalva/jwt-go"

// AuthForm class
type AuthForm struct{}

// CustomClaims object
type CustomClaims struct {
	Role string `json:"role" binding:"required"`
	jwt.StandardClaims
}

// Login object
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register object
type Register struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Type            string `json:"type" binding:"required"`
}
