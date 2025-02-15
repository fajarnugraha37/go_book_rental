package basev2

import (
	"backend/internal/database"
	"backend/internal/database/repo/filter"
	"backend/pkg/helper"
	"context"

	"github.com/uptrace/bun"
)

type repoDQLImpl[TModel any] struct {
	db *bun.DB
}

// FindOne implements RepoDQL.
func (r *repoDQLImpl[TModel]) FindOne(ctx context.Context, predicate *filter.Predicate) (*TModel, error) {
	result, query := database.UowSelect[TModel](ctx, r.db)
	err := query.
		Model(result).
		Apply(predicate.ToQuery()...).
		Limit(1).
		Scan(ctx)

	return helper.ReturnTuple(result, err)
}

// FindMany implements RepoDQL.
func (r *repoDQLImpl[TModel]) FindMany(ctx context.Context, predicate *filter.Predicate) ([]TModel, error) {
	result, query := database.UowSelect[[]TModel](ctx, r.db)
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
		return []TModel{}, err
	}

	return *result, nil
}

// FindMany implements RepoDQL.
func (r *repoDQLImpl[TModel]) FindPageable(ctx context.Context, predicate *filter.Predicate) (filter.Page[TModel], error) {
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
func (r *repoDQLImpl[TModel]) Count(ctx context.Context, predicate *filter.Predicate) (int, error) {
	_, query := database.UowSelect[[]TModel](ctx, r.db)
	count, err := query.
		Apply(predicate.ToQuery()...).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
