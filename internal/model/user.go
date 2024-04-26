package model

import "github.com/dgrijalva/jwt-go"

// UserForCreate ..
type UserForCreate struct {
	Name     string
	Email    string
	Password string
}

// UserClaims ..
type UserClaims struct {
	jwt.StandardClaims
	UserID   string `json:"userID"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}
