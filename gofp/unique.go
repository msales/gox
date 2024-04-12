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
