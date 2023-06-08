package gofp

// KeyBuilder is a type for a function that maps a value of type T to a comparable value of type C.
// This can be used for constructing a comparable key from an object, allowing it to be used in comparison operations.
type KeyBuilder[T any, C comparable] func(T) C

// SymDiffWithKeyBuilder calculates the symmetric difference between two slices.
// It uses a KeyBuilder function to map values in the slices to a comparable key, which is used for the comparison.
// The function returns a new slice containing the elements that are in slice1 or slice2 but not in both.
//
// Notes:
//   - This function doesn't return an ordered response.
//   - The KeyBuilder function is used to identify unique elements in the slices for comparison.
//     Therefore, it's important that the KeyBuilder function produces unique keys for unique elements.
//     If the keys are not unique, the function may not provide accurate results.
//     For example, consider two slices of a struct with attributes 'Name' and 'Age'. If the KeyBuilder function
//     uses only the 'Name' attribute for creating the key, two elements with the same 'Name' but different 'Age'
//     will be considered identical, which may not be the intended behavior.
//
// Example:
//
//	slice1 := []Person{
//		{Name: "Alice", Age: 25},
//		{Name: "Bob", Age: 30},
//		{Name: "Alice", Age: 40},
//	}
//	slice2 := []Person{
//		{Name: "Charlie", Age: 35},
//	}
//	kb := func(p Person) string {
//		return p.Name
//	}
//	result := SymDiffWithKeyBuilder(slice1, slice2, kb)
//
// The 'result' will be:
//
//	[]Person{
//			{Name: "Alice", Age: 25},
//			{Name: "Bob", Age: 30},
//			{Name: "Charlie", Age: 35},
//	}
//
// Note that for the 'Alice' element, the 'Age' attribute of 25 was used and the 40 was omitted.
func SymDiffWithKeyBuilder[T any, C comparable](slice1, slice2 []T, kb KeyBuilder[T, C]) []T {
	checks1, checks2 := make(Set[C]), make(Set[C])
	valueByKey := make(map[C]T)
	for _, v := range slice1 {
		k := kb(v)
		checks1.Add(k)
		valueByKey[k] = v
	}

	for _, v := range slice2 {
		checks2.Add(kb(v))
		k := kb(v)
		if _, ok := valueByKey[k]; !ok {
			valueByKey[k] = v
		}
	}

	var result []T
	for k := range checks1 {
		if !checks2.Has(k) {
			result = append(result, valueByKey[k])
		}
	}
	for k := range checks2 {
		if !checks1.Has(k) {
			result = append(result, valueByKey[k])
		}
	}

	return result
}

// DiffLeftWithKeyBuilder calculates the difference between two slices, i.e., elements that are in the first slice, but not in the second.
// It uses a KeyBuilder function to map values in the slices to a comparable key, which is used for the comparison.
// The function returns a new slice containing the elements that are in slice1 but not in slice2.
//
// Notes:
//   - This function does not retain the order of elements in the original slices. The order of elements in the resulting slice depends on the order of iteration over slice1.
//   - The KeyBuilder function is used to identify unique elements in the slices for comparison.
//     Therefore, it's important that the KeyBuilder function produces unique keys for unique elements.
//     If the keys are not unique, the function may not provide accurate results.
//     For example, consider two slices of a struct with attributes 'Name' and 'Age'. If the KeyBuilder function
//     uses only the 'Name' attribute for creating the key, two elements with the same 'Name' but different 'Age'
//     will be considered identical, which may not be the intended behavior.
//
// Example:
//
//	slice1 := []Person{
//		{Name: "Alice", Age: 25},
//		{Name: "Bob", Age: 30},
//		{Name: "Alice", Age: 40},
//		{Name: "Charlie", Age: 35},
//	}
//	slice2 := []Person{
//		{Name: "Charlie", Age: 20},
//	}
//	kb := func(p Person) string {
//		return p.Name
//	}
//	result := DiffLeftWithKeyBuilder(slice1, slice2, kb)
//
// The 'result' will be:
//
//	[]Person{
//		{Name: "Bob", Age: 30},
//		{Name: "Alice", Age: 25},
//	}
//
// Note that for the 'Alice' element, the 'Age' attribute of 25 was used and the 40 was omitted.
func DiffLeftWithKeyBuilder[T any, C comparable](slice1, slice2 []T, kb KeyBuilder[T, C]) []T {
	checks := make(Set[C])
	for _, v := range slice2 {
		checks.Add(kb(v))
	}

	outer := make([]T, 0)
	used := make(Set[C])
	for _, v := range slice1 {
		k := kb(v)
		if !checks.Has(k) && !used.Has(k) {
			outer = append(outer, v)
			used.Add(k)
		}
	}

	return outer
}

// DiffRightWithKeyBuilder calculates the difference between two slices, i.e., elements that are in the second slice, but not in the first.
// It uses a KeyBuilder function to map values in the slices to a comparable key, which is used for the comparison.
// The function returns a new slice containing the elements that are in slice2 but not in slice1.
//
// Notes:
//   - This function does not retain the order of elements in the original slices. The order of elements in the resulting slice depends on the order of iteration over slice2.
//   - The KeyBuilder function is used to identify unique elements in the slices for comparison.
//     Therefore, it's important that the KeyBuilder function produces unique keys for unique elements.
//     If the keys are not unique, the function may not provide accurate results.
//     For example, consider two slices of a struct with attributes 'Name' and 'Age'. If the KeyBuilder function
//     uses only the 'Name' attribute for creating the key, two elements with the same 'Name' but different 'Age'
//     will be considered identical, which may not be the intended behavior.
//
// Example:
//
//	slice1 := []Person{
//		{Name: "Charlie", Age: 20},
//	}
//	slice2 := []Person{
//		{Name: "Alice", Age: 25},
//		{Name: "Bob", Age: 30},
//		{Name: "Alice", Age: 40},
//		{Name: "Charlie", Age: 35},
//	}
//	kb := func(p Person) string {
//		return p.Name
//	}
//	result := DiffRightWithKeyBuilder(slice1, slice2, kb)
//
// The 'result' will be:
//
//	[]Person{
//		{Name: "Bob", Age: 30},
//		{Name: "Alice", Age: 25},
//	}
//
// Note that for the 'Alice' element, the 'Age' attribute of 25 was used and the 40 was omitted.
func DiffRightWithKeyBuilder[T any, C comparable](slice1, slice2 []T, kb KeyBuilder[T, C]) []T {
	return DiffLeftWithKeyBuilder(slice2, slice1, kb)
}
