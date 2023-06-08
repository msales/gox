package gofp

// SymDiff returns symmetric difference of 2 slices.
// The function returns a new slice containing the elements that are in slice1 or slice2 but not in both.
// https://en.wikipedia.org/wiki/Symmetric_difference
func SymDiff[T comparable](slice1, slice2 []T) []T {
	checks1 := make(map[T]bool)
	checks2 := make(map[T]bool)

	// Track unique elements in slice1 and slice2
	for _, v := range slice1 {
		checks1[v] = true
	}
	for _, v := range slice2 {
		checks2[v] = true
	}

	outer := make([]T, 0)

	// Add elements in slice1 but not in slice2
	for v := range checks1 {
		if !checks2[v] {
			outer = append(outer, v)
		}
	}

	// Add elements in slice2 but not in slice1
	for v := range checks2 {
		if !checks1[v] {
			outer = append(outer, v)
		}
	}

	return outer
}

// DiffLeft returns left diff of 2 slices.
// The function returns a new slice containing the elements that are in slice1 but not in slice2.
func DiffLeft[T comparable](slice1, slice2 []T) []T {
	checks := make(map[T]bool)
	for _, v := range slice2 {
		checks[v] = true
	}

	var outer []T
	used := make(map[T]bool)
	for _, v := range slice1 {
		if !checks[v] && !used[v] {
			used[v] = true
			outer = append(outer, v)
		}
	}

	return outer
}

// DiffRight returns right diff of 2 slices.
// The function returns a new slice containing the elements that are in slice2 but not in slice1.
func DiffRight[T comparable](slice1, slice2 []T) []T {
	return DiffLeft(slice2, slice1)
}
