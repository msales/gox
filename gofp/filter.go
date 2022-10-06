package gofp

// Filter apply a function to all elements on an array. Returns elements that are not filtered.
func Filter[T any](list []T, predicate func(T) bool) (res []T) {
	res = make([]T, 0, len(list))
	for i := range list {
		if !predicate(list[i]) {
			res = append(res, list[i])
		}
	}

	return
}
