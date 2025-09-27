package user

import (
	"context"
	"main/internal/entity"
	"main/internal/services/auth"
	"main/internal/services/user"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data auth.SignUp) (entity.User, error) {
	var detail entity.User

	detail.FullName = data.FullName
	detail.Email = data.Email
	detail.Password = &data.Password
	detail.Region = data.Region
	detail.City = data.City
	detail.Status = false

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var detail entity.User
	err := r.NewSelect().Model(&detail).Where("email = ?", email).Scan(ctx)
	return detail, err
}

func (r Repository) Update(ctx context.Context, data user.Update) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("id = ?", data.Id).Scan(ctx)
	if data.FullName != nil {
		detail.FullName = *data.FullName
	}
	if data.Password != nil {
		detail.Password = data.Password
	}
	if data.Email != nil {
		detail.Email = *data.Email
	}
	if data.Region != nil {
		detail.Region = *data.Region
	}
	if data.City != nil {
		detail.City = *data.City
	}
	if data.Avatar != nil {
		detail.Avatar = *data.Avatar
	}
	_, err = r.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)
	return detail, err
}

func (r Repository) UpdatePassword(ctx context.Context, email, password string) error {

	_, err := r.NewUpdate().
		Model(&entity.User{}).
		Set("password = ?", password).
		Where("email = ?", email).
		Exec(ctx)

	return err
}
