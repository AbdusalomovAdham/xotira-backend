package gallery

import (
	"context"
	"main/internal/entity"
	"mime/multipart"
)

type Repository interface {
	Create(ctx context.Context, imagePath string, id int) error
	GetAll(ctx context.Context, id int) ([]entity.GalleryImg, error)
	Delete(ctx context.Context, id int) error
}

type File interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	Delete(ctx context.Context, url string) error
}
