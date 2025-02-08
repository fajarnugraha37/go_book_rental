package basev2

import (
	"backend/internal/database/repo/filter"
	"context"

	"github.com/uptrace/bun"
)

type RepoDQLImpl[TModel any] struct {
	db *bun.DB
}

// FindOne implements RepoDQL.
func (r *RepoDQLImpl[TModel]) FindOne(ctx context.Context, predicate *filter.Predicate) (*TModel, error) {
	result := new(TModel)
	err := r.db.NewSelect().
		Model(result).
		Apply(predicate.ToQuery()...).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindMany implements RepoDQL.
func (r *RepoDQLImpl[TModel]) FindMany(ctx context.Context, predicate *filter.Predicate) ([]TModel, error) {
	var result []TModel
	query := r.db.NewSelect().
		Model(&result)

	if predicate.Pageable != nil {
		query = predicate.Pageable.Append(query)
	}
	if predicate.Sortable != nil {
		query = filter.SortablesAppend(*predicate.Sortable, query)
	}

	err := query.
		Apply(predicate.ToQuery()...).
		Scan(ctx)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindMany implements RepoDQL.
func (r *RepoDQLImpl[TModel]) FindPageable(ctx context.Context, predicate *filter.Predicate) (filter.Page[TModel], error) {
	items, err := r.FindMany(ctx, predicate)
	if err != nil {
		return filter.Page[TModel]{}, err
	}
	count, err := r.Count(ctx, predicate)
	if err != nil {
		return filter.Page[TModel]{}, err
	}

	return filter.NewPage(items, uint(count), predicate.Pageable.Page, predicate.Pageable.PageSize), nil
}

// Count implements RepoDQL.
func (r *RepoDQLImpl[TModel]) Count(ctx context.Context, predicate *filter.Predicate) (int, error) {
	model := new(TModel)
	count, err := r.db.NewSelect().
		Model(model).
		Apply(predicate.ToQuery()...).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
