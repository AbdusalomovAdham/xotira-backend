package region

import (
	"context"
	"main/internal/entity"
	"main/internal/services/region"
)

type UseCase struct {
	region Region
}

func NewUseCase(region Region) *UseCase {
	return &UseCase{region: region}
}

func (uc UseCase) GetAll(ctx context.Context, filter region.Filter, order, lang string) ([]entity.Region, int, error) {
	return uc.region.GetAll(ctx, filter, order, lang)
}
