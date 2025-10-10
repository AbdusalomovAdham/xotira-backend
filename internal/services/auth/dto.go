package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	FullName   string  `json:"full_name"`
	Email      string  `json:"email" `
	Password   string  `json:"password" `
	RegionId   int     `json:"region_id"`
	DistrictId int     `json:"district_id"`
	Avatar     *string `json:"avatar"`
}

type ForgotPsw struct {
	Email string
}
type GenerateToken struct {
	Id   int
	Role int
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
	ID    int
	Email string
	Role  int
	jwt.RegisteredClaims
}
