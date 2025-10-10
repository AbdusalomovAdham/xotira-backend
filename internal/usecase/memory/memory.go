package memory

import (
	"context"
	"errors"
	"main/internal/entity"
	"main/internal/services/memory"
	"mime/multipart"
	"os"
)

type UseCase struct {
	memory Memory
	auth   Auth
	file   File
}

func NewUseCase(memory Memory, auth Auth, file File) *UseCase {
	return &UseCase{memory: memory, auth: auth, file: file}
}

func (uc UseCase) CreateMemoryInCabinet(ctx context.Context, data memory.Create, authHeader string) (entity.Memory, error) {
	token, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return entity.Memory{}, err
	}

	data.CreatedBy = token.Id
	detail, err := uc.memory.Create(ctx, data)
	return detail, err
}

func (uc UseCase) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	return uc.file.Upload(ctx, file, folder)
}

func (uc UseCase) MultipleUpload(ctx context.Context, file []*multipart.FileHeader, folder string) ([]string, error) {
	return uc.file.MultipleUpload(ctx, file, folder)
}

func (uc UseCase) GetListInCabinet(ctx context.Context, authHeader string) ([]entity.Memory, error) {
	token, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return nil, err
	}

	return uc.memory.GetListInCabinet(ctx, token.Id)
}

func (uc UseCase) DeleteInCabinet(ctx context.Context, id int) error {
	return uc.memory.DeleteInCabinet(ctx, id)
}

func (uc UseCase) GetById(ctx context.Context, id int) (memory.GetById, error) {
	return uc.memory.GetById(ctx, id)
}

func (uc UseCase) UpdateInCabinet(ctx context.Context, data memory.Update, id int) (entity.Memory, error) {
	oldUser, err := uc.memory.GetById(ctx, id)
	if err != nil {
		return entity.Memory{}, err
	}

	if data.Avatar != nil && oldUser.Avatar != *data.Avatar {
		if err := uc.file.Delete(ctx, oldUser.Avatar); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return entity.Memory{}, err
			}
		}
	}

	updated, err := uc.memory.UpdateInCabinet(ctx, data, id)
	if err != nil {
		return entity.Memory{}, err
	}

	return updated, nil
}

func (uc UseCase) DeleteFile(ctx context.Context, file string) error {
	return uc.file.Delete(ctx, file)
}
