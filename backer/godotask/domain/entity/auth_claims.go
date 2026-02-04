package entity

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type AuthClaims struct {
	UserID   uint
	Username string
	Role     string
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
