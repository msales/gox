package slicex

import "strconv"

// Float64ToInterface converts float64 input to interface slice.
func Float64ToInterface(i []float64) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Float64ToString converts float64 input to string slice.
func Float64ToString(i []float64, precision int) []string {
	o := make([]string, len(i))
	for k, v := range i {
		o[k] = strconv.FormatFloat(v, 'f', precision, 64)
	}

	return o
}

// Float64ToInt converts float64 input to int slice.
func Float64ToInt(i []float64) []int {
	o := make([]int, len(i))
	for k, v := range i {
		o[k] = int(v)
	}

	return o
}

// Float64ToInt32 converts float64 input to int32 slice.
func Float64ToInt32(i []float64) []int32 {
	o := make([]int32, len(i))
	for k, v := range i {
		o[k] = int32(v)
	}

	return o
}

// Float64ToInt64 converts float64 input to int64 slice.
func Float64ToInt64(i []float64) []int64 {
	o := make([]int64, len(i))
	for k, v := range i {
		o[k] = int64(v)
	}

	return o
}

// Float64ToFloat32 converts float64 input to float32 slice.
func Float64ToFloat32(i []float64) []float32 {
	o := make([]float32, len(i))
	for k, v := range i {
		o[k] = float32(v)
	}

	return o
}

// ContainsFloat64 checks if all elements from needles list are in the haystack.
func ContainsFloat64(haystack []float64, needles ...float64) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

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

// ContainsAtLeastOneFloat64 checks if at least one element from needles list is in the haystack.
func ContainsAtLeastOneFloat64(haystack []float64, needles ...float64) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[float64]struct{}, len(needles))
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
