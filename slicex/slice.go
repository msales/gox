package slicex

import "reflect"

// Contains checks if any element exists in slice.
// It panics if x is not a slice.
func Contains(x interface{}, contains func(i int) bool) bool {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice {
		panic(&reflect.ValueError{Method: "Contains", Kind: v.Kind()})
	}

	for i := 0; i < v.Len(); i++ {
		if contains(i) {
			return true
		}
	}

	return false
}
