package memory

import (
	"github.com/lib/pq"
	"github.com/uptrace/bun"
)

type Create struct {
	Avatar         string   `form:"avatar"`
	FullName       string   `form:"full_name"`
	BirthPlace     string   `form:"birth_place"`
	BirthDate      string   `form:"birth_date"`
	FamilyMemberId int      `form:"family_member_id"`
	BioHeadline    string   `form:"bio_headline"`
	Bio            string   `form:"bio"`
	RegionId       int      `form:"region_id"`
	DistrictId     int      `form:"district_id"`
	CemeteryId     int      `form:"cemetery_id"`
	DeathPlace     string   `form:"death_place"`
	DeathDate      string   `form:"death_date"`
	DeathCauseId   int      `form:"death_cause_id"`
	MemoryStatus   bool     `form:"memory_status"`
	SocialId       []int    `form:"social_id"`
	Images         []string `form:"images"`
	Videos         []string `form:"videos"`
	Audio          []string `form:"audio"`
	CreatedBy      int
}

type GetById struct {
	ID               int            `json:"id"`
	Avatar           string         `json:"avatar"`
	FullName         string         `json:"full_name"`
	BirthPlace       string         `json:"birth_place"`
	BirthDate        string         `json:"birth_date"`
	FamilyMemberId   int            `json:"family_member_id"`
	FamilyMemberName string         `json:"family_member_name"`
	BioHeadline      string         `json:"bio_headline"`
	Bio              string         `json:"bio"`
	RegionName       string         `json:"region_name"`
	RegionId         int            `json:"region_id"`
	DistrictId       int            `json:"district_id"`
	DistrictName     string         `json:"district_name"`
	CemeteryName     string         `json:"cemetery_name"`
	CemeteryId       int            `json:"cemetery_id"`
	DeathPlace       string         `json:"death_place"`
	DeathDate        string         `json:"death_date"`
	DeathCauseId     int            `json:"death_cause"`
	MemoryStatus     bool           `json:"memory_status"`
	SocialId         pq.Int64Array  `json:"social_id" bun:"social_id,array"`
	Images           pq.StringArray `bun:"images,array" json:"images"`
	Videos           pq.StringArray `json:"videos"`
	Audio            pq.StringArray `json:"audio"`
}

type Detail struct {
	ID           int    `bun:"id"`
	FullName     string `bun:"full_name"`
	BirthPlace   string `bun:"birth_place"`
	BirthDate    string `bun:"birth_date"`
	BioHeadline  string `bun:"bio_headline"`
	Bio          string `bun:"bio"`
	DeathPlace   string `bun:"death_place"`
	DeathDate    string `bun:"death_date"`
	MemoryStatus bool   `bun:"memory_status"`
	DeathCauseId int    `bun:"death_cause_id"`

	ImagesRaw string `bun:"images"`
	VideosRaw string `bun:"videos"`
	AudioRaw  string `bun:"audio"`
	SocialRaw string `bun:"social_id"`

	RegionName       string `bun:"region_name"`
	DistrictName     string `bun:"district_name"`
	CemeteryName     string `bun:"cemetery_name"`
	FamilyMemberName string `bun:"family_member_name"`
}

type List struct {
	ID           int    `json:"id"`
	Avatar       string `json:"avatar"`
	FullName     string `json:"full_name"`
	FamilyMember string `json:"family_member"`
	DeathDate    string `json:"death_date"`
	DeathOfCause string `json:"death_cause"`
}

type Update struct {
	bun.BaseModel `bun:"table:memories"`

	ID           int `bun:"id,pk"`
	Avatar       *string
	FullName     *string `form:"full_name"`
	BirthPlace   *string `form:"birth_place"`
	BirthDate    *string `form:"birth_date"`
	BioHeadline  *string `form:"bio_headline"`
	Bio          *string `form:"bio"`
	DeathPlace   *string `form:"death_place"`
	DeathDate    *string `form:"death_date"`
	MemoryStatus *bool   `form:"memory_status"`
	DeathCauseId *int    `form:"death_cause_id"`

	Images        []string  `form:"images"`
	ImageExisting *[]string `form:"images_existing"`
	DeleteImg     *[]string `form:"delete_img"`

	Videos         []string  `form:"videos"`
	VideosExisting *[]string `form:"videos_existing"`
	DeleteVideo    *[]string `form:"delete_video"`

	Audio         []string  `form:"audio"`
	AudioExisting *[]string `form:"audio_existing"`
	DeleteAudio   *[]string `form:"delete_audio"`

	SocialId []int `form:"social_id"`

	RegionId       *int `form:"region_id"`
	DistrictId     *int `form:"district_id"`
	CemeteryId     *int `form:"cemetery_id"`
	FamilyMemberId *int `form:"family_member_id"`
}
