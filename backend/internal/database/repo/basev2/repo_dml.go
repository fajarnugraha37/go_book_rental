package basev2

import (
	"context"

	"github.com/uptrace/bun"
)

type RepoDML[TModel any] interface {
	Insert(ctx context.Context, model *TModel) error
	Update(ctx context.Context, model *TModel) error
	Delete(ctx context.Context, id string) (*TModel, error)
	Restore(ctx context.Context, id string) (*TModel, error)

	InsertBulk(ctx context.Context, model *[]TModel) error
	UpdateBulk(ctx context.Context, model *[]TModel) error

	// PartialUpdate(ctx context.Context, model *TResult) error
	// HardDelete(ctx context.Context, model *TResult) error
}

var _ RepoDML[any] = (*RepoDMLImpl[any])(nil)

func NewRepoDML[TModel any](db *bun.DB) *RepoDMLImpl[TModel] {
	return &RepoDMLImpl[TModel]{
		db: db,
	}
}
