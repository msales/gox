package slicex

import "strconv"

// Int64ToInterface converts int64 input to interface slice.
func Int64ToInterface(i []int64) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Int64ToString converts int64 input to string slice.
func Int64ToString(i []int64) []string {
	o := make([]string, len(i))
	for k, v := range i {
		o[k] = strconv.FormatInt(v, 10)
	}

	return o
}

// Int64ToInt converts int64 input to int slice.
func Int64ToInt(i []int64) []int {
	o := make([]int, len(i))
	for k, v := range i {
		o[k] = int(v)
	}

	return o
}

// Int64ToInt32 converts int64 input to int32 slice.
func Int64ToInt32(i []int64) []int32 {
	o := make([]int32, len(i))
	for k, v := range i {
		o[k] = int32(v)
	}

	return o
}

// Int64ToFloat32 converts int64 input to float32 slice.
func Int64ToFloat32(i []int64) []float32 {
	o := make([]float32, len(i))
	for k, v := range i {
		o[k] = float32(v)
	}

	return o
}

// Int64ToFloat64 converts int64 input to float64 slice.
func Int64ToFloat64(i []int64) []float64 {
	o := make([]float64, len(i))
	for k, v := range i {
		o[k] = float64(v)
	}

	return o
}
