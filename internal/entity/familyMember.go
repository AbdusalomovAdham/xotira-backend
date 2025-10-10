package entity

import "github.com/uptrace/bun"

type FamilyMember struct {
	bun.BaseModel `bun:"table:family_members"`
	ID            int    `json:"id" bun:"id,pk,autoincrement"`
	Name          string `json:"name" bun:"name"`
}
