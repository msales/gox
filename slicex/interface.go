package slicex

import "fmt"

// InterfaceToString converts interface input to string slice.
// Returns error on failed type assertion.
func InterfaceToString(i []interface{}) ([]string, error) {
	o := make([]string, len(i))
	for k, v := range i {
		val, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("slicex: value is not a string, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToInt converts interface input to primitive int slice.
// Returns error on failed type assertion.
func InterfaceToInt(i []interface{}) ([]int, error) {
	o := make([]int, len(i))
	for k, v := range i {
		val, ok := v.(int)
		if !ok {
			return nil, fmt.Errorf("slicex: value is not an int, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToInt32 converts interface input to int32 slice.
// Returns error on failed type assertion.
func InterfaceToInt32(i []interface{}) ([]int32, error) {
	o := make([]int32, len(i))
	for k, v := range i {
		val, ok := v.(int32)
		if !ok {
			return nil, fmt.Errorf("slicex: value is not an int32, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToInt64 converts interface input to int64 slice.
// Returns error on failed type assertion.
func InterfaceToInt64(i []interface{}) ([]int64, error) {
	o := make([]int64, len(i))
	for k, v := range i {
		val, ok := v.(int64)
		if !ok {
			return nil, fmt.Errorf("slicex: value is not an int64, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToFloat32 converts interface input to float32 slice.
// Returns error on failed type assertion.
func InterfaceToFloat32(i []interface{}) ([]float32, error) {
	o := make([]float32, len(i))
	for k, v := range i {
		val, ok := v.(float32)
		if !ok {
			return nil, fmt.Errorf("slicex: value is not an float32, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToFloat64 converts interface input to float64 slice.
// Returns error on failed type assertion.
func InterfaceToFloat64(i []interface{}) ([]float64, error) {
	o := make([]float64, len(i))
	for k, v := range i {
		val, ok := v.(float64)
		if !ok {
			return nil, fmt.Errorf("slicex: value is not an float64, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}
