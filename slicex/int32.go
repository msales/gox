package slicex

import "strconv"

// Int32ToInterface converts int32 input to interface slice.
func Int32ToInterface(i []int32) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Int32ToString converts int32 input to string slice.
func Int32ToString(i []int32) []string {
	o := make([]string, len(i))
	for k, v := range i {
		o[k] = strconv.FormatInt(int64(v), 10)
	}

	return o
}

// Int32ToInt converts int32 input to int slice.
func Int32ToInt(i []int32) []int {
	o := make([]int, len(i))
	for k, v  := range i {
		o[k] = int(v)
	}

	return o
}

// Int32ToInt64 converts int32 input to int64 slice.
func Int32ToInt64(i []int32) []int64 {
	o := make([]int64, len(i))
	for k, v  := range i {
		o[k] = int64(v)
	}

	return o
}

// Int32ToFloat32 converts int32 input to float32 slice.
func Int32ToFloat32(i []int32) []float32 {
	o := make([]float32, len(i))
	for k, v  := range i {
		o[k] = float32(v)
	}

	return o
}

// Int32ToFloat64 converts int32 input to float64 slice.
func Int32ToFloat64(i []int32) []float64 {
	o := make([]float64, len(i))
	for k, v  := range i {
		o[k] = float64(v)
	}

	return o
}
