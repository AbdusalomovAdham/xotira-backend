package user

import (
	"time"
)

type Create struct {
	FullName   string `form:"full_name"`
	Password   string `form:"password"`
	RegionId   int    `form:"region_id"`
	DistrictId int    `form:"district_id"`
	Email      string `form:"email"`
	Avatar     string
	Role       *int `form:"role"`
	CreatedBy  int
	UpdatedBy  int
}

type Update struct {
	Id         *int    `json:"id" form:"id"`
	FullName   *string `json:"full_name" form:"full_name"`
	Password   *string `json:"password" form:"password"`
	RegionId   *int    `json:"region_id" form:"region_id"`
	DistrictId *int    `json:"district_id" form:"district_id"`
	Avatar     *string
	Email      *string   `json:"email" form:"email"`
	UpdatedBy  int       `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	Status     *bool     `json:"status" form:"status"`
	Role       *int      `json:"role" form:"role"`
}

type Filter struct {
	Limit  *int
	Offset *int
	Role   *int
	Status *bool
}

type UserWithLocation struct {
	FullName     string `json:"full_name"`
	RegionName   string `json:"region_name"`
	DistrictName string `json:"district_name"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	Role         int    `json:"role"`
}

type UpdateCabinet struct {
	Fullname   *string `json:"full_name" form:"full_name"`
	Email      *string `json:"email" form:"email"`
	RegionId   *int    `json:"region_id" form:"region_id"`
	DistrictId *int    `json:"district_id" form:"district_id"`
	Avatar     *string
}
