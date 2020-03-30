package slicex

import (
	"golang.org/x/xerrors"
)

// StringToInterface converts primitive typed slice to interface slice.
func StringToInterface(i []string) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// IntToInterface converts primitive typed slice to interface slice.
func IntToInterface(i []int) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Int32ToInterface converts primitive typed slice to interface slice.
func Int32ToInterface(i []int32) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Int64ToInterface converts primitive typed slice to interface slice.
func Int64ToInterface(i []int64) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Float32ToInterface converts primitive typed slice to interface slice.
func Float32ToInterface(i []float32) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// Float64ToInterface converts primitive typed slice to interface slice.
func Float64ToInterface(i []float64) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

// InterfaceToString converts interface slice to primitive typed slice.
// Returns error on failed type assertion.
func InterfaceToString(i []interface{}) ([]string, error) {
	o := make([]string, len(i))
	for k, v := range i {
		val, ok := v.(string)
		if !ok {
			return nil, xerrors.Errorf("slicex: value is not a string, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToInt converts interface slice to primitive typed slice.
// Returns error on failed type assertion.
func InterfaceToInt(i []interface{}) ([]int, error) {
	o := make([]int, len(i))
	for k, v := range i {
		val, ok := v.(int)
		if !ok {
			return nil, xerrors.Errorf("slicex: value is not an int, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToInt32 converts interface slice to primitive typed slice.
// Returns error on failed type assertion.
func InterfaceToInt32(i []interface{}) ([]int32, error) {
	o := make([]int32, len(i))
	for k, v := range i {
		val, ok := v.(int32)
		if !ok {
			return nil, xerrors.Errorf("slicex: value is not an int32, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToInt64 converts interface slice to primitive typed slice.
// Returns error on failed type assertion.
func InterfaceToInt64(i []interface{}) ([]int64, error) {
	o := make([]int64, len(i))
	for k, v := range i {
		val, ok := v.(int64)
		if !ok {
			return nil, xerrors.Errorf("slicex: value is not an int64, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToFloat32 converts interface slice to primitive typed slice.
// Returns error on failed type assertion.
func InterfaceToFloat32(i []interface{}) ([]float32, error) {
	o := make([]float32, len(i))
	for k, v := range i {
		val, ok := v.(float32)
		if !ok {
			return nil, xerrors.Errorf("slicex: value is not an float32, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// InterfaceToFloat64 converts interface slice to primitive typed slice.
// Returns error on failed type assertion.
func InterfaceToFloat64(i []interface{}) ([]float64, error) {
	o := make([]float64, len(i))
	for k, v := range i {
		val, ok := v.(float64)
		if !ok {
			return nil, xerrors.Errorf("slicex: value is not an float64, received %T instead", v)
		}

		o[k] = val
	}

	return o, nil
}

// ContainsInt checks if all elements from `contains` list are on the `slice` list
func ContainsInt(slice []int, contains ...int) bool {
	toCheck := len(contains)

	for _, s := range slice {
		for _, c := range contains {
			if s != c {
				continue
			}

			toCheck--
			if toCheck == 0 {
				return true
			}
		}
	}

	return false
}
