package region

import (
	"context"
	"main/internal/entity"
	"main/internal/services/region"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetAll(ctx context.Context, filter region.Filter, order, lang string) ([]entity.Region, int, error) {
	var list []entity.Region

	q := r.NewSelect().Model(&list)

	if filter.Offset != nil {
		q.Offset(*filter.Offset)
	}

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}

	if order == "asc" {
		q.Order("id asc")
	} else {
		q.Order("id desc")
	}

	q.Column("id").ColumnExpr("name ->> ? AS name", lang)
	
	count, err := q.ScanAndCount(ctx)
	return list, count, err
}
