package slicex

import "strconv"

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
