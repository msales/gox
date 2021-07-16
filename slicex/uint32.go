package slicex

// ContainsUint32 checks if all elements from needles list are in the haystack.
func ContainsUint32(haystack []uint32, needles ...uint32) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[uint32]struct{}, len(needles))
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

// ContainsAtLeastOneUint32 checks if at least one element from needles list is in the haystack.
func ContainsAtLeastOneUint32(haystack []uint32, needles ...uint32) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[uint32]struct{}, len(needles))
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

// Uint32ToInterface converts uint32 input to interface slice.
func Uint32ToInterface(i []uint32) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}
