package gallery

import (
	"context"
	"main/internal/entity"
	"mime/multipart"
)

type Service struct {
	repo Repository
	file File
}

func NewService(repo Repository, file File) *Service {
	return &Service{repo: repo, file: file}
}

func (s Service) Create(ctx context.Context, imagePath string, id int) error {
	return s.repo.Create(ctx, imagePath, id)
}
func (s Service) GetAll(ctx context.Context, id int) ([]entity.GalleryImg, error) {
	return s.repo.GetAll(ctx, id)
}

func (s Service) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	return s.file.Upload(ctx, file, folder)
}

func (s Service) DeleteFile(ctx context.Context, url string) error {
	return s.file.Delete(ctx, url)
}

func (s Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
