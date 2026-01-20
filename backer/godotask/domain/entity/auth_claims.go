package entity

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type AuthClaims struct {
	UserID   uint
	Username string
	Role     string
	Email    string `json:"email"`
	jwt.StandardClaims
}
