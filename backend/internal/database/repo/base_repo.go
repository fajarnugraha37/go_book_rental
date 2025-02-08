package repo

import (
	"context"

	"github.com/uptrace/bun"
)

type baseRepo[TFilter, TResult any] interface {
	FindOne(ctx context.Context, filter TFilter, sortable *Sortable) (*TResult, error)
	FindAll(ctx context.Context, filter TFilter, pageable *Pageable, sortable *Sortable) (*[]TResult, error)
	Insert(ctx context.Context, model *TResult) error
	Update(ctx context.Context, model *TResult) error
	Delete(ctx context.Context, id string) (*TResult, error)
	Restore(ctx context.Context, id string) (*TResult, error)

	// PartialUpdate(ctx context.Context, model *TResult) error
	// HardDelete(ctx context.Context, model *TResult) error
}

type baseRepoImpl[TFilter, TResult any] struct {
	db *bun.DB
}

func newBaseRepo[TFilter, TResult any](db *bun.DB) *baseRepoImpl[TFilter, TResult] {
	var _ baseRepo[any, any] = (*baseRepoImpl[any, any])(nil)
	return &baseRepoImpl[TFilter, TResult]{
		db: db,
	}
}

// FindOne implements BaseRepo.
func (b *baseRepoImpl[TFilter, TResult]) FindOne(ctx context.Context, filter TFilter, sortable *Sortable) (*TResult, error) {
	result := new(TResult)
	q := b.db.NewSelect().
		Model(result).
		Limit(1)
	if sortable != nil {
		q = sortable.Append(q)
	}

	err := q.ApplyQueryBuilder(filterToQueryBuilder(filter)).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindAll implements BaseRepo.
func (b *baseRepoImpl[TFilter, TResult]) FindAll(ctx context.Context, filter TFilter, pageable *Pageable, sortable *Sortable) (*[]TResult, error) {
	result := new([]TResult)
	q := b.db.NewSelect().
		Model(result)
	if pageable != nil {
		q = pageable.Append(q)
	}
	if sortable != nil {
		q = sortable.Append(q)
	}

	err := q.ApplyQueryBuilder(filterToQueryBuilder(filter)).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert implements Insert.
func (b *baseRepoImpl[TFilter, TResult]) Insert(ctx context.Context, model *TResult) error {
	_, err := b.db.NewInsert().
		Model(model).
		Returning("*").
		Exec(ctx)

	return err
}

// Update implements Update.
func (b *baseRepoImpl[TFilter, TResult]) Update(ctx context.Context, model *TResult) error {
	_, err := b.db.NewUpdate().
		Model(model).
		WherePK().
		Returning("*").
		Exec(ctx)

	return err
}

// Update implements Delete.
func (b *baseRepoImpl[TFilter, TResult]) Delete(ctx context.Context, id string) (*TResult, error) {
	result := new(TResult)
	_, err := b.db.NewUpdate().
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

// Update implements Restore.
func (b *baseRepoImpl[TFilter, TResult]) Restore(ctx context.Context, id string) (*TResult, error) {
	result := new(TResult)
	_, err := b.db.NewUpdate().
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
