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
		arguments  []any = []any{bun.Ident(field)}
	)
	switch comparator {
	case EQ:
		expression = " ? = ? "
		arguments = append(arguments, value)
	case NEQ:
		expression = " ? != ? "
		arguments = append(arguments, value)
	case LT:
		expression = " ? < ? "
	case LTE:
		expression = " ? <= ? "
		arguments = append(arguments, value)
	case GT:
		expression = " ? > ? "
		arguments = append(arguments, value)
	case GTE:
		expression = " ? >= ? "
		arguments = append(arguments, value)
	case LIKE:
		expression = " ? LIKE ? "
		arguments = append(arguments, value)
	case NLIKE:
		expression = " ? NOT LIKE ? "
		arguments = append(arguments, value)
	case IN:
		expression = " ? IN (?) "
		arguments = append(arguments, bun.In(value))
	case NIN:
		expression = " ? NOT IN (?) "
		arguments = append(arguments, bun.In(value))
	case BETWEEN:
		expression = " ? BETWEEN ? AND ? "
		arguments = append(arguments, value, secondValue)
	case NBETWEEN:
		expression = " ? NOT BETWEEN ? AND ? "
		arguments = append(arguments, value, secondValue)
	default:
		panic(fmt.Errorf("unsupported comparator %+v", comparator))
	}

	return expression, arguments
}
