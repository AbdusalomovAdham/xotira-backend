package memory

import (
	"context"
	"main/internal/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, data Create) (entity.Memory, error) {
	return s.repo.Create(ctx, data)
}

func (s Service) GetListInCabinet(ctx context.Context, id int) ([]entity.Memory, error) {
	return s.repo.GetListInCabinet(ctx, id)
}

func (s Service) DeleteInCabinet(ctx context.Context, id int) error {
	return s.repo.DeleteInCabinet(ctx, id)
}

func (s Service) GetById(ctx context.Context, id int) (GetById, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) UpdateInCabinet(ctx context.Context, data Update, id int) (entity.Memory, error) {
	return s.repo.UpdateInCabinet(ctx, data, id)
}
