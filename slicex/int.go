package slicex

import (
	"strconv"
)

// IntToInterface converts int input to interface slice.
func IntToInterface(i []int) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// IntToString converts int input to string slice.
func IntToString(i []int) []string {
	o := make([]string, len(i))
	for k, v := range i {
		o[k] = strconv.Itoa(v)
	}

	return o
}

// IntToInt32 converts int input to int32 slice.
func IntToInt32(i []int) []int32 {
	o := make([]int32, len(i))
	for k, v := range i {
		o[k] = int32(v)
	}

	return o
}

// IntToInt64 converts int input to int64 slice.
func IntToInt64(i []int) []int64 {
	o := make([]int64, len(i))
	for k, v := range i {
		o[k] = int64(v)
	}

	return o
}

// IntToFloat32 converts int input to float32 slice.
func IntToFloat32(i []int) []float32 {
	o := make([]float32, len(i))
	for k, v := range i {
		o[k] = float32(v)
	}

	return o
}

// IntToFloat64 converts int input to float64 slice.
func IntToFloat64(i []int) []float64 {
	o := make([]float64, len(i))
	for k, v := range i {
		o[k] = float64(v)
	}

	return o
}

// ContainsInt checks if all elements from needles list are in the haystack.
func ContainsInt(haystack []int, needles ...int) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

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

// IntOuter returns outersection of 2 slices.
func IntOuter(slice1, slice2 []int) []int {
	checks := make(map[int]int8, len(slice1)+len(slice2))
	idx := make([]int, 0, len(slice1)+len(slice2))
	for _, v := range slice1 {
		checks[v]++
		idx = append(idx, v)
	}
	for _, v := range slice2 {
		checks[v]++
		if checks[v] == 1 {
			idx = append(idx, v)
		}
	}

	outer := make([]int, 0, len(slice1)+len(slice2))
	for id := range idx {
		if checks[id] == 1 {
			outer = append(outer, id)
		}
	}

	return outer
}

// IntOuterLeft returns left outersection of 2 slices.
func IntOuterLeft(slice1, slice2 []int) []int {
	checks := make(map[int]int8, len(slice1))
	idx := make([]int, 0, len(slice1))
	for _, v := range slice1 {
		checks[v]++
		idx = append(idx, v)
	}
	for _, v := range slice2 {
		checks[v]++
	}

	outer := make([]int, 0, len(slice1))
	for _, id := range idx {
		if checks[id] == 1 {
			outer = append(outer, id)
		}
	}

	return outer
}

// IntOuterRight returns right outersection of 2 slices.
func IntOuterRight(slice1, slice2 []int) []int {
	checks := make(map[int]int8, len(slice2))
	idx := make([]int, 0, len(slice2))
	for _, v := range slice2 {
		checks[v]++
		idx = append(idx, v)
	}
	for _, v := range slice1 {
		checks[v]++
	}

	outer := make([]int, 0, len(slice2))
	for _, id := range idx {
		if checks[id] == 1 {
			outer = append(outer, id)
		}
	}

	return outer
}
