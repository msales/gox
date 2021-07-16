package slicex

import "strconv"

// Float32ToInterface converts float32 input to interface slice.
func Float32ToInterface(i []float32) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Float32ToString converts float32 input to string slice.
func Float32ToString(i []float32, precision int) []string {
	o := make([]string, len(i))
	for k, v := range i {
		o[k] = strconv.FormatFloat(float64(v), 'f', precision, 64)
	}

	return o
}

// Float32ToInt converts float32 input to int slice.
func Float32ToInt(i []float32) []int {
	o := make([]int, len(i))
	for k, v := range i {
		o[k] = int(v)
	}

	return o
}

// Float32ToInt32 converts float32 input to int32 slice.
func Float32ToInt32(i []float32) []int32 {
	o := make([]int32, len(i))
	for k, v := range i {
		o[k] = int32(v)
	}

	return o
}

// Float32ToInt64 converts float32 input to int64 slice.
func Float32ToInt64(i []float32) []int64 {
	o := make([]int64, len(i))
	for k, v := range i {
		o[k] = int64(v)
	}

	return o
}

// Float32ToFloat64 converts float32 input to float64 slice.
func Float32ToFloat64(i []float32) []float64 {
	o := make([]float64, len(i))
	for k, v := range i {
		o[k] = float64(v)
	}

	return o
}

// ContainsFloat32 checks if all elements from needles list are in the haystack.
func ContainsFloat32(haystack []float32, needles ...float32) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

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

// ContainsAtLeastOneFloat32 checks if at least one element from needles list is in the haystack.
func ContainsAtLeastOneFloat32(haystack []float32, needles ...float32) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[float32]struct{}, len(needles))
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
