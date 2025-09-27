package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignUp struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Region   string `json:"region" binding:"required"`
	City     string `json:"city" binding:"required"`
}

type ForgotPsw struct {
	Email string
}
type GenerateToken struct {
	Email string
	Role  int
}

type CheckCode struct {
	Token string
	Code  string
}

type ResetData struct {
	Email string
	Code  string
}

type UpdatePsw struct {
	Token    string
	Password string
}

type ResendCode struct {
	Token string
}

type Claims struct {
	Email string
	Role  int
	jwt.RegisteredClaims
}
