package gofp

// KeyBuilder is a type for a function that maps a value of type T to a comparable value of type C.
// This can be used for constructing a comparable key from an object, allowing it to be used in comparison operations.
type KeyBuilder[T any, C comparable] func(T) C

// SymDiffWithKeyBuilder is a function that calculates the symmetric difference between two slices.
// It uses a KeyBuilder function to map values in the slices to a comparable key, which is used for the comparison.
// The function returns a new slice containing the elements that are in slice1 or slice2 but not in both.
func SymDiffWithKeyBuilder[T any, C comparable](slice1, slice2 []T, kb KeyBuilder[T, C]) []T {
	checks := make(map[C]int8, len(slice1)+len(slice2))
	idx := make([]T, 0, len(slice1)+len(slice2))
	for _, v := range slice1 {
		checks[kb(v)]++
		idx = append(idx, v)
	}

	for _, v := range slice2 {
		checks[kb(v)]++
		if checks[kb(v)] == 1 {
			idx = append(idx, v)
		}
	}

	outer := make([]T, 0, len(slice1)+len(slice2))
	for _, id := range idx {
		if checks[kb(id)] == 1 {
			outer = append(outer, id)
		}
	}

	return outer
}

// DiffLeftWithKeyBuilder is a function that calculates the difference between two slices.
// It uses a KeyBuilder function to map values in the slices to a comparable key, which is used for the comparison.
// The function returns a new slice containing the elements that are in slice1 but not in slice2.
func DiffLeftWithKeyBuilder[T any, C comparable](slice1, slice2 []T, kb KeyBuilder[T, C]) []T {
	checks := make(map[C]int8, len(slice1))
	for _, v := range slice1 {
		checks[kb(v)]++
	}
	for _, v := range slice2 {
		checks[kb(v)]++
	}

	outer := make([]T, 0, len(slice1))
	for _, id := range slice1 {
		if checks[kb(id)] == 1 {
			outer = append(outer, id)
		}
	}

	return outer
}

// DiffRightWithKeyBuilder is a function that calculates the difference between two slices.
// It uses a KeyBuilder function to map values in the slices to a comparable key, which is used for the comparison.
// The function returns a new slice containing the elements that are in slice2 but not in slice1.
// It's essentially a reverse version of DiffLeftWithKeyBuilder.
func DiffRightWithKeyBuilder[T any, C comparable](slice1, slice2 []T, kb KeyBuilder[T, C]) []T {
	return DiffLeftWithKeyBuilder(slice2, slice1, kb)
}
