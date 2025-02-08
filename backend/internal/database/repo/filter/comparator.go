package filter

import (
	"fmt"
	"strings"

	"github.com/uptrace/bun"
)

type ComparatorType uint

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
	BETWEEN
	NBETWEEN
)

var comparatorString = []string{
	"EQ",
	"NEQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"LIKE",
	"NLIKE",
	"IN",
	"NIN",
	"BETWEEN",
	"NBETWEEN",
}

func (comparator ComparatorType) String() string {
	return comparatorString[comparator]
}

func CompartorEnum(value string) ComparatorType {
	for i, v := range comparatorString {
		if strings.EqualFold(v, value) {
			return ComparatorType(i)
		}
	}

	panic(fmt.Errorf("invalid query comparator: %+v", value))
}

func (comparator ComparatorType) toExpression(field string, value, secondValue any) (string, []any) {
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
	case BETWEEN:
		expression = " BETWEEN ? AND ? "
		arguments = []any{value, secondValue}
	case NBETWEEN:
		expression = " NOT BETWEEN ? AND ? "
		arguments = []any{value, secondValue}
	default:
		panic(fmt.Errorf("unsupported comparator %+v", comparator))
	}
	expression = field + expression

	return expression, arguments
}
