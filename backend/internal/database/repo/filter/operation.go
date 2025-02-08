package filter

import (
	"fmt"
	"strings"
)

type OperationType uint

const (
	OR = iota
	AND
)

var operationString = []string{
	"OR",
	"AND",
}

func (operation OperationType) IsOr() bool {
	return operation == OR
}

func (operation OperationType) String() string {
	return operationString[operation]
}

func OperationEnum(value string) OperationType {
	for i, v := range operationString {
		if strings.EqualFold(v, value) {
			return OperationType(i)
		}
	}

	panic(fmt.Errorf("invalid query operation: %+v", value))
}
