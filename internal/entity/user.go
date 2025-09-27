package entity

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id       int64   `json:"id" bun:"id,pk,autoincrement"`
	FullName string  `json:"full_name" bun:"full_name"`
	Email    string  `json:"email" bun:"email"`
	Password *string `json:"password" bun:"password"`
	Region   string  `json:"region" bun:"region"`
	City     string  `json:"city" bun:"city"`
	Status   bool    `json:"status" bun:"status,default:false"`
	Role     int     `json:"role" bun:"role,default:3"`
	Avatar   string  `json:"avatar" bun:"avatar"`
}

//Role

//1-super amdin
//2-admin
//3-user
