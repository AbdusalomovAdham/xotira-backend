package districts

import (
	"main/internal/entity"

	"golang.org/x/net/context"
)

type Repository interface {
	GetByRegionId(ctx context.Context, retion_id int, lang string) ([]entity.District, error)
}
