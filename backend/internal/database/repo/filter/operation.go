package filter

import "strings"

type OperationType string

func (operation OperationType) IsOr() bool {
	return strings.ToLower(string(operation)) == "or"
}

func (operation OperationType) ToOperation() string {
	if operation.IsOr() {
		return " OR "
	} else {
		return " AND "
	}
}
