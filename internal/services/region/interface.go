package region

import (
	"context"
	"main/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter, order string, lang string) ([]entity.Region, int, error)
}
