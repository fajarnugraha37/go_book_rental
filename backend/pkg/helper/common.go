package helper

func ToPtr[T any](t T) *T {
	return &t
}

func Must[T any](runner func() (T, error)) T {
	result, err := runner()
	if err != nil {
		panic(err)
	}

	return result
}

func ReturnTuple[T any](result *T, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	return result, nil
}
