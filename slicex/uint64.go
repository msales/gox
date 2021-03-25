package slicex

// ContainsUint64 checks if all elements from needles list are in the haystack.
func ContainsUint64(haystack []uint64, needles ...uint64) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[uint64]struct{}, len(needles))
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

// Uint64ToInterface converts uint64 input to interface slice.
func Uint64ToInterface(i []uint64) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}
