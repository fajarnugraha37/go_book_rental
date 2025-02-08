package filter

import (
	"fmt"

	"github.com/uptrace/bun"
)

const (
	EQ = iota
	NEQ
	LT
	LTE
	GT
	GTE
	LIKE
	NLIKE
	IN
	NIN
)

type ComparatorType uint

func (comparator ComparatorType) ToExpression(field string, value any) (string, []any) {
	var (
		expression string
		arguments  []any = []any{value}
	)
	switch comparator {
	case EQ:
		expression = " = ? "
	case NEQ:
		expression = " != ? "
	case LT:
		expression = " < ? "
	case LTE:
		expression = " <= ? "
	case GT:
		expression = " > ? "
	case GTE:
		expression = " >= ? "
	case LIKE:
		expression = " LIKE ? "
	case NLIKE:
		expression = " NOT LIKE ? "
	case IN:
		expression = " IN (?) "
		arguments = []any{bun.In(value)}
	case NIN:
		expression = " NOT IN (?) "
		arguments = []any{bun.In(value)}
	default:
		panic(fmt.Errorf("unsupported comparator %+v", comparator))
	}
	expression = field + expression

	return expression, arguments
}
