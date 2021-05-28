package slicex

import "reflect"

// Contains checks if any element exists in slice.
// It panics if x is not a slice.
func Contains(x interface{}, contains func(i int) bool) bool {
	v := reflect.ValueOf(x)
	validateReflection("Contains", reflect.Slice, v)
	for i := 0; i < v.Len(); i++ {
		if contains(i) {
			return true
		}
	}

	return false
}

// Filter filters out elements from slice.
// It panics if x is not a pointer to a slice.
func Filter(x interface{}, filter func(i int) bool) {
	ptr := reflect.ValueOf(x)
	validateReflection("Filter", reflect.Ptr, ptr)

	v := ptr.Elem()
	validateReflection("Filter", reflect.Slice, v)

	for i := 0; i < v.Len(); i++ {
		if filter(i) {
			v.Set(reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())))
		}
	}
}

func validateReflection(method string, kind reflect.Kind, v reflect.Value) {
	if v.Kind() != kind {
		panic(&reflect.ValueError{Method: method, Kind: v.Kind()})
	}
}
