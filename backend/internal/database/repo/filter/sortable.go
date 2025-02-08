package filter

import (
	"strings"

	"github.com/uptrace/bun"
)

type DirectionType uint8

const (
	ASC DirectionType = iota
	DESC
)

type Sortable struct {
	SortBy    string
	Direction DirectionType
}

func (s Sortable) ToDirection() string {
	direction := "ASC"
	if s.Direction == 1 {
		direction = "DESC"
	}

	return direction
}

func SortablesAppend(s []Sortable, q *bun.SelectQuery) *bun.SelectQuery {
	var (
		expressions []string
		args        []interface{}
	)

	for i := 0; i < len(s); i++ {
		sortable := s[i]
		expressions = append(expressions, "? ?")
		args = append(args, bun.Safe(sortable.SortBy), bun.Safe(sortable.ToDirection()))
	}
	q.OrderExpr(strings.Join(expressions, ", ")+" NULLS LAST", args...)

	return q
}
