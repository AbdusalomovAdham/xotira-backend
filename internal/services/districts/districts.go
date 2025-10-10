package districts

import (
	"main/internal/entity"

	"golang.org/x/net/context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetByRegionId(ctx context.Context, region_id int, lang string) ([]entity.District, error) {
	return s.repo.GetByRegionId(ctx, region_id, lang)
}
