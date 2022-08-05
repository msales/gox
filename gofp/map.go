package gofp

// Map apply a function to all elements on an array
func Map[T, V any](list []T, fn func(T) V) (res []V) {
	for i := range list {
		res = append(res, fn(list[i]))
	}

	return
}
