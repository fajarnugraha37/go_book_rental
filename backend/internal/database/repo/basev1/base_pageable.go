package basev1

import "github.com/uptrace/bun"

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

const (
	ASC uint8 = iota
	DESC
)

type Sortable struct {
	SortBy    string
	Direction uint8
}

func (s Sortable) Append(q *bun.SelectQuery) *bun.SelectQuery {
	if s.Direction != 0 || s.Direction != 1 || s.SortBy == "" {
		return q
	}

	direction := "ASC"
	if s.Direction == 1 {
		direction = "DESC"
	}
	q.OrderExpr("? ? NULLS LAST",
		bun.Safe(s.SortBy),
		bun.Safe(direction),
	)

	return q
}
