package region

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

func (s Service) GetAll(ctx context.Context, filter Filter, order, lang string) ([]entity.Region, int, error) {
	return s.repo.GetAll(ctx, filter, order, lang)
}
