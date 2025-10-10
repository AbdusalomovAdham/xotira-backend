package user

import (
	"context"
	"errors"
	"main/internal/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, data Create) (entity.User, error) {

	if data.Email == "" {
		return entity.User{}, errors.New("email is required")
	}
	if data.Role == nil {
		return entity.User{}, errors.New("role is required")
	}
	if data.FullName == "" {
		return entity.User{}, errors.New("full name is required")
	}
	return s.repo.Create(ctx, data)
}

func (s Service) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s Service) IsValidToken(ctx context.Context, authHeader entity.User) (entity.User, error) {
	return s.IsValidToken(ctx, authHeader)
}

func (s Service) GetAll(ctx context.Context, filter Filter, order string) ([]entity.User, int, error) {
	return s.repo.GetAll(ctx, filter, order)
}

func (s Service) GetById(ctx context.Context, id int) (entity.User, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) Update(ctx context.Context, data Update, created_by int) (entity.User, error) {
	return s.repo.Update(ctx, data, created_by)
}

func (s Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s Service) GetByEmailWithLocation(ctx context.Context, id int, lang string) (UserWithLocation, error) {
	return s.repo.GetByEmailWithLocation(ctx, id, lang)
}

func (s Service) UpdateCabinet(ctx context.Context, data UpdateCabinet, id int) (entity.User, error) {
	return s.repo.UpdateCabinet(ctx, data, id)
}
