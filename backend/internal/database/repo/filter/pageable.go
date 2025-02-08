package filter

import (
	"math"

	"github.com/uptrace/bun"
)

type Pageable struct {
	Page     uint
	PageSize uint
}

func (p Pageable) Append(q *bun.SelectQuery) *bun.SelectQuery {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 20
	}
	q.Limit(int(p.PageSize)).
		Offset(int((p.Page - 1) * p.PageSize))
	return q
}

type Page[TModel any] struct {
	Items      []TModel
	TotalItems uint
	Page       uint
	NextPage   uint
	PageSize   uint
	TotalPages uint
}

func NewPage[TModel any](items []TModel, count, currentPage, pageSize uint) Page[TModel] {
	return Page[TModel]{
		Items:      items,
		TotalItems: count,
		Page:       currentPage,
		NextPage:   (currentPage + 1),
		PageSize:   pageSize,
		TotalPages: uint(math.Ceil(float64(count) / float64(pageSize))),
	}
}

func (p Page[any]) HasNext() bool {
	count := uint(len(p.Items))
	if count < p.PageSize || p.Page == p.TotalPages {
		return false
	}

	return true
}

func (p Page[any]) NextPageable() Pageable {
	return Pageable{
		Page:     p.NextPage,
		PageSize: p.PageSize,
	}
}
