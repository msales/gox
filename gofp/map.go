package gofp

// Map apply a function to all elements on an array
func Map[T, V any](list []T, fn func(T) V) (res []V) {
	res = make([]V, len(list))
	for i := range list {
		res[i] = fn(list[i])
	}

	return
}

// MapMultiple apply a function to all elements on an array.
// Applied function should return multiple elements out of single list element.
func MapMultiple[T, V any](list []T, fn func(T) []V) (res []V) {
	for i := range list {
		res = append(res, fn(list[i])...)
	}

	return
}
