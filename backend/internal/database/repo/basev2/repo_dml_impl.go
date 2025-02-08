package basev2

import (
	"context"

	"github.com/uptrace/bun"
)

type RepoDMLImpl[TModel any] struct {
	db *bun.DB
}

// Insert implements RepoDML.
func (r *RepoDMLImpl[TModel]) Insert(ctx context.Context, model *TModel) error {
	_, err := r.db.NewInsert().
		Model(model).
		Returning("*").
		Exec(ctx)

	return err
}

// Update implements RepoDML.
func (r *RepoDMLImpl[TModel]) Update(ctx context.Context, model *TModel) error {
	_, err := r.db.NewUpdate().
		Model(model).
		WherePK().
		Returning("*").
		Exec(ctx)

	return err
}

// Delete implements RepoDML.
func (r *RepoDMLImpl[TModel]) Delete(ctx context.Context, id string) (*TModel, error) {
	result := new(TModel)
	_, err := r.db.NewUpdate().
		Model(result).
		Set("deleted_flag = true").
		Set("deleted_at = current_timestamp").
		Set("deleted_by = 'system'").
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Restore implements RepoDML.
func (r *RepoDMLImpl[TModel]) Restore(ctx context.Context, id string) (*TModel, error) {
	result := new(TModel)
	_, err := r.db.NewUpdate().
		Model(result).
		Set("deleted_flag = false").
		Set("updated_at = current_timestamp").
		Set("updated_by = 'system'").
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
