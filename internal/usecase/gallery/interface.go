package gallery

import (
	"context"
	"main/internal/entity"
	"mime/multipart"
)

type Gallery interface {
	Create(ctx context.Context, imagePath string, id int) error
	GetAll(ctx context.Context, id int) ([]entity.GalleryImg, error)
	Delete(ctx context.Context, id int) error
}

type File interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	Delete(ctx context.Context, url string) error
}

type Auth interface {
	IsValidToken(ctx context.Context, token string) (entity.User, error)
}
