package gofp

// Contains returns if exist at least one element that satisfy the predicate.
func Contains[T any](list []T, predicate func(T) bool) (res bool) {
	for i := range list {
		if predicate(list[i]) {
			return true
		}
	}

	return false
}

// ContainsAll checks if all elements from needles list are in the haystack.
func ContainsAll[T comparable](haystack []T, needles ...T) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[T]struct{}, len(needles))
	for _, n := range needles {
		checks[n] = struct{}{}
	}

	for _, h := range haystack {
		_, ok := checks[h]
		if ok {
			delete(checks, h)

			if len(checks) == 0 {
				return true
			}
		}
	}

	return false
}

// ContainsAtLeastOne checks if at least one element from needles list is in the haystack.
func ContainsAtLeastOne[T comparable](haystack []T, needles ...T) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[T]struct{}, len(needles))
	for _, n := range needles {
		checks[n] = struct{}{}
	}

	for _, h := range haystack {
		_, ok := checks[h]
		if ok {
			return true
		}
	}

	return false
}
