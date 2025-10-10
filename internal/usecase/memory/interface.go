package memory

import (
	"context"
	"main/internal/entity"
	"main/internal/services/memory"
	"mime/multipart"
)

type Memory interface {
	Create(ctx context.Context, data memory.Create) (entity.Memory, error)
	GetListInCabinet(ctx context.Context, id int) ([]entity.Memory, error)
	DeleteInCabinet(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (memory.GetById, error)
	UpdateInCabinet(ctx context.Context, data memory.Update, id int) (entity.Memory, error)
}

type Auth interface {
	IsValidToken(ctx context.Context, tokenStr string) (entity.User, error)
}

type File interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	MultipleUpload(ctx context.Context, file []*multipart.FileHeader, folder string) ([]string, error)
	Delete(ctx context.Context, url string) error
}
