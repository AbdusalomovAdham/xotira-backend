package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type District struct {
	bun.BaseModel `bun:"table:districts"`
	Id            int       `json:"id" bun:"id,pk,autoincrement"`
	Name          string    `json:"name" bun:"name"`
	RegionId      int       `json:"region_id" bun:"region_id"`
	IsDelete      bool      `json:"is_deleted" bun:"is_deleted"`
	Uid           int       `json:"uid" bun:"uid"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	DeletedAt     time.Time `json:"deleted_at" bun:"deleted_at"`
	DeletedBy     int       `json:"deleted_by" bun:"deleted_by"`
	CreatedBy     int       `json:"created_by" bun:"created_by"`
}
