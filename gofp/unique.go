package gofp

// IsUnique checks if all elements in a slice are unique.
func IsUnique[T comparable](el []T) bool {
	seen := map[T]bool{}
	for _, t := range el {
		_, ok := seen[t]
		if ok {
			return false
		}

		seen[t] = true
	}

	return true
}

// IsUniqueWithKeyBuilder checks if all elements in a slice after being transformed by the KeyBuilder function.
func IsUniqueWithKeyBuilder[T any, C comparable](el []T, kb KeyBuilder[T, C]) bool {
	seen := make(Set[C])
	for _, t := range el {
		k := kb(t)
		if seen.Has(k) {
			return false
		}

		seen.Add(k)
	}

	return true
}
