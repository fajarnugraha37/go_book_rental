package helper

func ToPtr[T any](t T) *T {
	return &t
}
