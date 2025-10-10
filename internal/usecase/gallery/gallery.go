package gallery

import (
	"context"
	"main/internal/entity"
	"mime/multipart"
)

type UseCase struct {
	gallery Gallery
	file    File
	auth    Auth
}

func NewUseCase(gallery Gallery, file File, auth Auth) *UseCase {
	return &UseCase{gallery: gallery, file: file, auth: auth}
}

func (uc UseCase) Create(ctx context.Context, imagePath string, authHeader string) error {
	token, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}
	id := token.Id
	return uc.gallery.Create(ctx, imagePath, id)
}

func (uc UseCase) GetAll(ctx context.Context, authHeader string) ([]entity.GalleryImg, error) {
	token, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return nil, err
	}
	id := token.Id
	return uc.gallery.GetAll(ctx, id)
}

func (uc UseCase) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	return uc.file.Upload(ctx, file, folder)
}

func (uc UseCase) Delete(ctx context.Context, id int, url string) error {
	if err := uc.file.Delete(ctx, url); err != nil {
		return err
	}
	return uc.gallery.Delete(ctx, id)
}
