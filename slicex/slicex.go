package slicex

import (
	"errors"
	"fmt"
)

func StringToInterface(i []string) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

func IntToInterface(i []int) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

func Int32ToInterface(i []int32) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

func Int64ToInterface(i []int64) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

func Float32ToInterface(i []float32) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

func Float64ToInterface(i []float64) []interface{} {
	o := make([]interface{}, len(i))
	for k, v := range i {
		o[k] = v
	}

	return o
}

func InterfaceToString(i []interface{}) ([]string, error) {
	o := make([]string, len(i))
	for k, v := range i {
		val, ok := v.(string)
		if !ok {
			return nil, errors.New(fmt.Sprintf("slicex: value is not a string, received %T instead", v))
		}

		o[k] = val
	}

	return o, nil
}

func InterfaceToInt(i []interface{}) ([]int, error) {
	o := make([]int, len(i))
	for k, v := range i {
		val, ok := v.(int)
		if !ok {
			return nil, errors.New(fmt.Sprintf("slicex: value is not an int, received %T instead", v))
		}

		o[k] = val
	}

	return o, nil
}

func InterfaceToInt32(i []interface{}) ([]int32, error) {
	o := make([]int32, len(i))
	for k, v := range i {
		val, ok := v.(int32)
		if !ok {
			return nil, errors.New(fmt.Sprintf("slicex: value is not an int32, received %T instead", v))
		}

		o[k] = val
	}

	return o, nil
}

func InterfaceToInt64(i []interface{}) ([]int64, error) {
	o := make([]int64, len(i))
	for k, v := range i {
		val, ok := v.(int64)
		if !ok {
			return nil, errors.New(fmt.Sprintf("slicex: value is not an int64, received %T instead", v))
		}

		o[k] = val
	}

	return o, nil
}

func InterfaceToFloat32(i []interface{}) ([]float32, error) {
	o := make([]float32, len(i))
	for k, v := range i {
		val, ok := v.(float32)
		if !ok {
			return nil, errors.New(fmt.Sprintf("slicex: value is not an float32, received %T instead", v))
		}

		o[k] = val
	}

	return o, nil
}

func InterfaceToFloat64(i []interface{}) ([]float64, error) {
	o := make([]float64, len(i))
	for k, v := range i {
		val, ok := v.(float64)
		if !ok {
			return nil, errors.New(fmt.Sprintf("slicex: value is not an float64, received %T instead", v))
		}

		o[k] = val
	}

	return o, nil
}