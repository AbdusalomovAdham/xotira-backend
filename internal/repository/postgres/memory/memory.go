package memory

import (
	"context"
	"fmt"
	"main/internal/entity"
	"main/internal/services/memory"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data memory.Create) (entity.Memory, error) {
	var detail entity.Memory

	detail.Avatar = data.Avatar
	detail.FullName = data.FullName
	detail.RegionId = data.RegionId
	detail.BirthDate = data.BirthDate
	detail.BirthPlace = data.BirthPlace
	detail.BioHeadline = data.BioHeadline
	detail.Bio = data.Bio
	detail.RegionId = data.RegionId
	detail.DistrictId = data.DistrictId
	detail.CemeteryId = data.CemeteryId
	detail.DeathPlace = data.DeathPlace
	detail.DeathDate = data.DeathDate
	detail.DeathCauseId = data.DeathCauseId
	detail.SocialId = data.SocialId
	detail.Images = data.Images
	detail.Videos = data.Videos
	detail.Audio = data.Audio
	detail.MemoryStatus = data.MemoryStatus
	detail.CreatedBy = data.CreatedBy
	detail.FamilyMemberId = data.FamilyMemberId

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) GetListInCabinet(ctx context.Context, id int) ([]entity.Memory, error) {
	var list []entity.Memory

	err := r.NewSelect().
		Model(&list).
		Relation("FamilyMember").
		Where("created_by = ?", id).
		Scan(ctx)

	return list, err
}

func (r Repository) DeleteInCabinet(ctx context.Context, id int) error {
	_, err := r.NewDelete().
		Table("memories").
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (r *Repository) GetById(ctx context.Context, id int) (memory.GetById, error) {
	var m memory.GetById

	err := r.DB.NewSelect().
		Model((*entity.Memory)(nil)).
		ColumnExpr(`
			memory.id,
			memory.avatar,
			memory.full_name,
			memory.birth_place,
			memory.birth_date,
			memory.bio_headline,
			memory.bio,
			memory.region_id,
			memory.district_id,
			memory.cemetery_id,
			memory.death_place,
			memory.death_date,
			memory.death_cause_id,
			memory.memory_status,
			memory.social_id,
			memory.images,
			memory.videos,
			memory.audio,
			f.id AS family_member_id,
			f.name AS family_member_name,
			r.name ->> 'ru' AS region_name,
			d.name ->> 'ru' AS district_name,
			c.name AS cemetery_name
		`).
		Join("LEFT JOIN regions AS r ON r.id = memory.region_id").
		Join("LEFT JOIN districts AS d ON d.id = memory.district_id").
		Join("LEFT JOIN cemeteries AS c ON c.id = memory.cemetery_id").
		Join("LEFT JOIN family_members AS f ON f.id = memory.family_member_id").
		Where("memory.id = ?", id).
		Scan(ctx, &m)

	if err != nil {
		return m, err
	}

	return m, nil
}

func (r Repository) UpdateInCabinet(ctx context.Context, data memory.Update, id int) (entity.Memory, error) {
	var detail entity.Memory

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return entity.Memory{}, err
	}

	if data.Avatar != nil {
		detail.Avatar = *data.Avatar
	}

	if data.FullName != nil {
		detail.FullName = *data.FullName
	}

	if data.BirthPlace != nil {
		detail.BirthPlace = *data.BirthPlace
	}

	if data.BirthDate != nil {
		detail.BirthDate = *data.BirthDate
	}

	if data.BioHeadline != nil {
		detail.BioHeadline = *data.BioHeadline
	}

	if data.Bio != nil {
		detail.Bio = *data.Bio
		fmt.Println("data hello repo", *data.Bio, detail.Bio)
	}

	if data.DeathPlace != nil {
		detail.DeathPlace = *data.DeathPlace
	}

	if data.DeathDate != nil {
		detail.DeathDate = *data.DeathDate
	}

	if data.MemoryStatus != nil {
		detail.MemoryStatus = *data.MemoryStatus
	}

	if data.DeathCauseId != nil {
		detail.DeathCauseId = *data.DeathCauseId
	}

	if len(data.Images) > 0 {
		detail.Images = data.Images
	}

	if len(data.Videos) > 0 {
		detail.Videos = data.Videos
	}

	if len(data.Audio) > 0 {
		detail.Audio = data.Audio
	}

	if len(data.SocialId) > 0 {
		detail.SocialId = data.SocialId
	}

	if data.RegionId != nil {
		detail.RegionId = *data.RegionId
	}

	if data.DistrictId != nil {
		detail.DistrictId = *data.DistrictId
	}

	if data.CemeteryId != nil {
		detail.CemeteryId = *data.CemeteryId
	}

	if data.FamilyMemberId != nil {
		detail.FamilyMemberId = *data.FamilyMemberId
	}

	_, err = r.NewUpdate().Model(&detail).Where("id = ?", id).Exec(ctx)

	return detail, nil
}
