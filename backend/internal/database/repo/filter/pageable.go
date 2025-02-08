package filter

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
