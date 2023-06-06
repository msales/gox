package gofp

// SymDiff returns symmetric difference of 2 slices.
// The function returns a new slice containing the elements that are in slice1 or slice2 but not in both.
// https://en.wikipedia.org/wiki/Symmetric_difference
func SymDiff[T comparable](slice1, slice2 []T) []T {
	checks := make(map[T]int8, len(slice1)+len(slice2))
	idx := make([]T, 0, len(slice1)+len(slice2))
	for _, v := range slice1 {
		checks[v]++
		idx = append(idx, v)
	}
	for _, v := range slice2 {
		checks[v]++
		if checks[v] == 1 {
			idx = append(idx, v)
		}
	}

	outer := make([]T, 0, len(slice1)+len(slice2))
	for _, id := range idx {
		if checks[id] == 1 {
			outer = append(outer, id)
		}
	}

	return outer
}

// DiffLeft returns left diff of 2 slices.
// The function returns a new slice containing the elements that are in slice1 but not in slice2.
func DiffLeft[T comparable](slice1, slice2 []T) []T {
	checks := make(map[T]int8, len(slice1))
	for _, v := range slice1 {
		checks[v]++
	}
	for _, v := range slice2 {
		checks[v]++
	}

	outer := make([]T, 0, len(slice1))
	for _, id := range slice1 {
		if checks[id] == 1 {
			outer = append(outer, id)
		}
	}

	return outer
}

// DiffRight returns right diff of 2 slices.
// The function returns a new slice containing the elements that are in slice2 but not in slice1.
func DiffRight[T comparable](slice1, slice2 []T) []T {
	return DiffLeft(slice2, slice1)
}
