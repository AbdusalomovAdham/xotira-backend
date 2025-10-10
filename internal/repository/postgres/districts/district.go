package districts

import (
	"context"
	"main/internal/entity"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetByRegionId(ctx context.Context, regionID int, lang string) ([]entity.District, error) {
	var list []entity.District

	err := r.NewSelect().
		Table("districts").
		Column("id").
		ColumnExpr("name ->> ? AS name", lang).
		Where("region_id = ?", regionID).
		Scan(ctx, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}
