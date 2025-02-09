package basev2

import (
	"backend/internal/database"
	"context"

	"github.com/uptrace/bun"
)

type RepoDMLImpl[TModel any] struct {
	db *bun.DB
}

// Insert implements RepoDML.
func (r *RepoDMLImpl[TModel]) Insert(ctx context.Context, model *TModel) error {
	query := database.UowInsert(ctx, r.db, model)
	_, err := query.
		Returning("*").
		Exec(ctx)

	return err
}

// InsertBulk implements RepoDML.
func (r *RepoDMLImpl[TModel]) InsertBulk(ctx context.Context, models *[]TModel) error {
	query := database.UowInsert(ctx, r.db, models)
	_, err := query.
		Returning("*").
		Exec(ctx)

	return err
}

// Update implements RepoDML.
func (r *RepoDMLImpl[TModel]) Update(ctx context.Context, model *TModel) error {
	query := database.UowUpdate(ctx, r.db, model)
	_, err := query.
		WherePK().
		Returning("*").
		Exec(ctx)

	return err
}

// UpdateBulk implements RepoDML.
func (r *RepoDMLImpl[TModel]) UpdateBulk(ctx context.Context, models *[]TModel) error {
	query := database.UowUpdate(ctx, r.db, models)
	_, err := query.
		WherePK().
		Returning("*").
		Exec(ctx)

	return err
}

// Delete implements RepoDML.
func (r *RepoDMLImpl[TModel]) Delete(ctx context.Context, id string) (*TModel, error) {
	model := new(TModel)
	query := database.UowUpdate(ctx, r.db, model)
	_, err := query.
		Set("deleted_flag = true").
		Set("deleted_at = current_timestamp").
		Set("deleted_by = 'system'").
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// Restore implements RepoDML.
func (r *RepoDMLImpl[TModel]) Restore(ctx context.Context, id string) (*TModel, error) {
	model := new(TModel)
	query := database.UowUpdate(ctx, r.db, model)
	_, err := query.
		Set("deleted_flag = false").
		Set("updated_at = current_timestamp").
		Set("updated_by = 'system'").
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}
