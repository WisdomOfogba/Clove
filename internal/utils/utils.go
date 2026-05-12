package utils

func AsPointer[T any](a T) *T {
	return &a
}

