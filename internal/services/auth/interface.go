package auth

import (
	"context"
	"main/internal/entity"
)

type User interface {
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}
