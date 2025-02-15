package basev2

import (
	"backend/internal/database"
	"backend/pkg/helper"
	"context"

	"github.com/uptrace/bun"
)

type repoDMLImpl[TModel any] struct {
	db *bun.DB
}

// Insert implements RepoDML.
func (r *repoDMLImpl[TModel]) Insert(ctx context.Context, model *TModel) error {
	query := database.UowInsert(ctx, r.db, model)
	_, err := query.
		Returning("*").
		Exec(ctx)

	return err
}

// InsertBulk implements RepoDML.
func (r *repoDMLImpl[TModel]) InsertBulk(ctx context.Context, models *[]TModel) error {
	query := database.UowInsert(ctx, r.db, models)
	_, err := query.
		Returning("*").
		Exec(ctx)

	return err
}

// Update implements RepoDML.
func (r *repoDMLImpl[TModel]) Update(ctx context.Context, model *TModel) error {
	query := database.UowUpdate(ctx, r.db, model)
	_, err := query.
		WherePK().
		Returning("*").
		Exec(ctx)

	return err
}

// UpdateBulk implements RepoDML.
func (r *repoDMLImpl[TModel]) UpdateBulk(ctx context.Context, models *[]TModel) error {
	query := database.UowUpdate(ctx, r.db, models)
	_, err := query.
		WherePK().
		Returning("*").
		Exec(ctx)

	return err
}

// Delete implements RepoDML.
func (r *repoDMLImpl[TModel]) Delete(ctx context.Context, id string) (*TModel, error) {
	model := new(TModel)
	query := database.UowUpdate(ctx, r.db, model)
	_, err := query.
		Set("deleted_flag = true").
		Set("deleted_at = current_timestamp").
		Set("deleted_by = 'system'").
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)

	return helper.ReturnTuple(model, err)
}

// Restore implements RepoDML.
func (r *repoDMLImpl[TModel]) Restore(ctx context.Context, id string) (*TModel, error) {
	model := new(TModel)
	query := database.UowUpdate(ctx, r.db, model)
	_, err := query.
		Set("deleted_flag = false").
		Set("updated_at = current_timestamp").
		Set("updated_by = 'system'").
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)

	return helper.ReturnTuple(model, err)
}
