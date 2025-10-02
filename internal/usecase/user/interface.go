package user

import (
	"context"
	"main/internal/entity"
	"main/internal/services/user"
	"mime/multipart"
)

type User interface {
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, data user.Create) (entity.User, error)
	Update(ctx context.Context, data user.Update, id int) (entity.User, error)
	GetAll(ctx context.Context, filer user.Filter, order string) ([]entity.User, int, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	Delete(ctx context.Context, id int) error
}

type Auth interface {
	HashPassword(password string) (string, error)
	IsValidToken(ctx context.Context, tokenStr string) (entity.User, error)
}

type File interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	Delete(ctx context.Context, url string) error
}
