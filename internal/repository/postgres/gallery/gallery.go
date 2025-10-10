package gallery

import (
	"context"
	"main/internal/entity"
	"time"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, imagePath string, id int) error {
	galleryImg := entity.GalleryImg{
		Image:      imagePath,
		CreatedBy:  id,
		CreatedBAt: time.Now(),
	}
	_, err := r.NewInsert().Model(&galleryImg).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetAll(ctx context.Context, id int) ([]entity.GalleryImg, error) {
	var list []entity.GalleryImg

	err := r.NewSelect().Model(&list).Where("created_by = ?", id).Order("created_at DESC").Scan(ctx)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r Repository) Delete(ctx context.Context, id int) error {
	_, err := r.NewDelete().Table("gallery").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
