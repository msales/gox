package slicex

// ContainsString checks if all elements from needles list are in the haystack.
func ContainsString(haystack []string, needles ...string) bool {
	checks := make(map[string]struct{}, len(needles))
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

// ContainsInt checks if all elements from needles list are in the haystack.
func ContainsInt(haystack []int, needles ...int) bool {
	checks := make(map[int]struct{}, len(needles))
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

// ContainsInt32 checks if all elements from needles list are in the haystack.
func ContainsInt32(haystack []int32, needles ...int32) bool {
	checks := make(map[int32]struct{}, len(needles))
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

// ContainsInt64 checks if all elements from needles list are in the haystack.
func ContainsInt64(haystack []int64, needles ...int64) bool {
	checks := make(map[int64]struct{}, len(needles))
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

// ContainsFloat32 checks if all elements from needles list are in the haystack.
func ContainsFloat32(haystack []float32, needles ...float32) bool {
	checks := make(map[float32]struct{}, len(needles))
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

// ContainsFloat64 checks if all elements from needles list are in the haystack.
func ContainsFloat64(haystack []float64, needles ...float64) bool {
	checks := make(map[float64]struct{}, len(needles))
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


