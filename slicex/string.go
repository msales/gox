package slicex

import (
	"fmt"
	"strconv"
)

// StringToInterface converts string input to interface slice.
func StringToInterface(i []string) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// StringToInt converts string input to int slice.
// Returns error on failed type assertion.
func StringToInt(i []string) ([]int, error) {
	o := make([]int, len(i))
	for k, v := range i {
		c, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("slicex: unable to convert string '%s' to int", v)
		}

		o[k] = c
	}

	return o, nil
}

// StringToInt32 converts string input to int32 slice.
// Returns error on failed type assertion.
func StringToInt32(i []string) ([]int32, error) {
	o := make([]int32, len(i))
	for k, v := range i {
		c, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("slicex: unable to convert string '%s' to int32", v)
		}

		o[k] = int32(c)
	}

	return o, nil
}

// StringToInt64 converts string input to int64 slice.
// Returns error on failed type assertion.
func StringToInt64(i []string) ([]int64, error) {
	o := make([]int64, len(i))
	for k, v := range i {
		c, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("slicex: unable to convert string '%s' to int64", v)
		}

		o[k] = c
	}

	return o, nil
}

// StringToInt64 converts string input to float32 slice.
// Returns error on failed type assertion.
func StringToFloat32(i []string) ([]float32, error) {
	o := make([]float32, len(i))
	for k, v := range i {
		c, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, fmt.Errorf("slicex: unable to convert string '%s' to float32", v)
		}

		o[k] = float32(c)
	}

	return o, nil
}

// StringToInt64 converts string input to float64 slice.
// Returns error on failed type assertion.
func StringToFloat64(i []string) ([]float64, error) {
	o := make([]float64, len(i))
	for k, v := range i {
		c, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("slicex: unable to convert string '%s' to float64", v)
		}

		o[k] = c
	}

	return o, nil
}
