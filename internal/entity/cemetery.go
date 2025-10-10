package entity

import "github.com/uptrace/bun"

type Cemetery struct {
	bun.BaseModel `bun:"table:cemeteries"`

	ID         int    `json:"id" bun:"id,pk, autoincrement"`
	Name       string `json:"cemetery" bun:"name"`
	DistrictId int    `bun:"district_id"`
}
