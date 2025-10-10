package districts

import (
	"context"
	"main/internal/entity"
)

type Districts interface {
	GetByRegionId(ctx context.Context, region_id int, lang string) ([]entity.District, error)
}
