package region

import (
	"context"
	"main/internal/entity"
	"main/internal/services/region"
)

type Region interface {
	GetAll(ctx context.Context, filter region.Filter, order, lang string) ([]entity.Region, int, error)
}
