package gofp

func Count[T any](list []T, predicate func(T) bool) (res int64) {
	for i := range list {
		if predicate(list[i]) {
			res++
		}
	}
	return
}
