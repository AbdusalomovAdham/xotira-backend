package memory

import (
	"context"
	"main/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, data Create) (entity.Memory, error)
	GetListInCabinet(ctx context.Context, id int) ([]entity.Memory, error)
	DeleteInCabinet(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (GetById, error)
	UpdateInCabinet(ctx context.Context, data Update, id int) (entity.Memory, error)
}
