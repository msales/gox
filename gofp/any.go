package gofp

// Any returns if exist at least one element that satisfy the predicate.
func Any[T any](list []T, predicate func(T) bool) (res bool) {
	for i := range list {
		if predicate(list[i]) {
			return true
		}
	}

	return false
}
