package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Memory struct {
	bun.BaseModel `bun:"table:memories"`

	ID             int       `json:"id" bun:"id,pk,autoincrement"`
	Avatar         string    `json:"avatar" bun:"avatar"`
	FullName       string    `json:"full_name" bun:"full_name"`
	BirthPlace     string    `json:"birth_place" bun:"birth_place"`
	BirthDate      string    `json:"birth_date" bun:"birth_date"`
	FamilyMemberId int       `json:"family_member_id" bun:"family_member_id"`
	BioHeadline    string    `json:"bio_headline" bun:"bio_headline"`
	Bio            string    `json:"bio" bun:"bio"`
	RegionId       int       `json:"region_id" bun:"region_id"`
	DistrictId     int       `json:"district_id" bun:"district_id"`
	CemeteryId     int       `json:"cemetery_id" bun:"cemetery_id"`
	DeathPlace     string    `json:"death_place" bun:"death_place"`
	DeathDate      string    `json:"death_date" bun:"death_date"`
	DeathCauseId   int       `json:"death_cause_id" bun:"death_cause_id"`
	MemoryStatus   bool      `json:"memory_status" bun:"memory_status"`
	SocialId       []int     `json:"social_id" bun:"social_id,array,nullzero"`
	Images         []string  `json:"images" bun:"images,array,nullzero"`
	Videos         []string  `json:"videos" bun:"videos,array,nullzero"`
	Audio          []string  `json:"audio" bun:"audio,array,nullzero"`
	CreatedBy      int       `json:"created_by" bun:"created_by"`
	CreatedAt      time.Time `json:"created_at" bun:"created_at"`

	Region       *Region       `bun:"rel:belongs-to,join:region_id=id"`
	District     *District     `bun:"rel:belongs-to,join:district_id=id"`
	Cemetery     *Cemetery     `bun:"rel:belongs-to,join:cemetery_id=id"`
	FamilyMember *FamilyMember `bun:"rel:belongs-to,join:family_member_id=id"`
}
