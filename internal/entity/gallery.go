package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type GalleryImg struct {
	bun.BaseModel `bun:"table:gallery"`
	ID            int       `json:"id" bun:"id,pk,autoincrement"`
	Image         string    `json:"image" bun:"image"`
	CreatedBy     int       `json:"created_by" bun:"created_by"`
	CreatedBAt    time.Time `json:"created_at" bun:"created_at"`
}
