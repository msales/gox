package gofp

// Merge merges passed slices
func Merge[T any](lists ...[]T) (res []T) {
	for _, list := range lists {
		res = append(res, list...)
	}

	return
}
