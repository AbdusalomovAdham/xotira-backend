package districts

import (
	"context"
	"main/internal/entity"
)

type UseCase struct {
	districts Districts
}

func NewUseCase(districts Districts) *UseCase {
	return &UseCase{districts: districts}
}

func (uc UseCase) GetByRegionId(ctx context.Context, region_id int, lang string) ([]entity.District, error) {
	return uc.districts.GetByRegionId(ctx, region_id, lang)
}
