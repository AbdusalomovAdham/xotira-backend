package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id         int       `json:"id" bun:"id,pk,autoincrement"`
	FullName   string    `json:"full_name" bun:"full_name"`
	Email      string    `json:"email" bun:"email"`
	Password   *string   `json:"password" bun:"password"`
	RegionId   int       `json:"region_id" bun:"region_id"`
	DistrictId int       `json:"district_id" bun:"district_id"`
	Status     bool      `json:"status" bun:"status,default:false"`
	Role       int       `json:"role" bun:"role,default:3"`
	Avatar     string    `json:"avatar" bun:"avatar"`
	CreatedBy  int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt  time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy  int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt  time.Time `json:"updated_at" bun:"updated_at,default:null"`
}

//Role

//1-super amdin
//2-admin
//3-user
