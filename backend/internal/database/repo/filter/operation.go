package filter

type OperationType uint

const (
	OR = iota
	AND
)

func (operation OperationType) IsOr() bool {
	return operation == OR
}

func (operation OperationType) ToOperation() string {
	if operation.IsOr() {
		return " OR "
	} else {
		return " AND "
	}
}
