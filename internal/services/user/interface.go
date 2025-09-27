package user

import (
	"context"
	"main/internal/entity"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, data entity.User) (entity.User, error)
	Update(ctx context.Context, data Update) (entity.User, error)
}
