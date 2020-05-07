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
	for k, v  := range i {
		o[k] = int(v)
	}

	return o
}

// Float64ToInt32 converts float64 input to int32 slice.
func Float64ToInt32(i []float64) []int32 {
	o := make([]int32, len(i))
	for k, v  := range i {
		o[k] = int32(v)
	}

	return o
}

// Float64ToInt64 converts float64 input to int64 slice.
func Float64ToInt64(i []float64) []int64 {
	o := make([]int64, len(i))
	for k, v  := range i {
		o[k] = int64(v)
	}

	return o
}

// Float64ToFloat32 converts float64 input to float32 slice.
func Float64ToFloat32(i []float64) []float32 {
	o := make([]float32, len(i))
	for k, v  := range i {
		o[k] = float32(v)
	}

	return o
}