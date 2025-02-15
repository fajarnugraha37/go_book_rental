package basev2

import (
	"backend/internal/database/repo/filter"
	"context"

	"github.com/uptrace/bun"
)

type RepoDQL[TModel any] interface {
	FindOne(ctx context.Context, predicate *filter.Predicate) (*TModel, error)
	FindMany(ctx context.Context, predicate *filter.Predicate) ([]TModel, error)
	FindPageable(ctx context.Context, predicate *filter.Predicate) (filter.Page[TModel], error)
	Count(ctx context.Context, predicate *filter.Predicate) (int, error)
}

var _ RepoDQL[any] = (*repoDQLImpl[any])(nil)

func NewRepoDQL[TModel any](db *bun.DB) RepoDQL[TModel] {
	return &repoDQLImpl[TModel]{
		db: db,
	}
}
