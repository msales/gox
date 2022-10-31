package gofp

import "fmt"

// Set[T] represents a list of unique 'T' elements.
// NOTE: only works with comparable types because the underlying implementation
// uses maps and map keys can only be 'comparables'.
// Should be builded using make. Example:
// integerSet:= make(Set[int]).
type Set[T comparable] map[T]struct{}

// Add adds a new element to the set.
func (s Set[T]) Add(val ...T) {
	for _, v := range val {
		s[v] = struct{}{}
	}
}

// Has verifies if an element exists in the set.
func (s Set[T]) Has(val T) bool {
	_, ok := s[val]
	return ok
}

// Remove removes an element from the set.
func (s Set[T]) Remove(val T) {
	delete(s, val)
}

// String implements Stringer
func (s Set[T]) String() string {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}

	return fmt.Sprintf("%v", keys)
}

// ToSlice convert the set in a slice.
func (s Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
