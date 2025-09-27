package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Region   string `json:"region"`
	City     string `json:"city"`
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

type Claims struct {
	Email string
	Role  int
	jwt.RegisteredClaims
}
