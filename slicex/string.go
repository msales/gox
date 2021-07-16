package slicex

import (
	"fmt"
	"strconv"
	"strings"
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

// StringToFloat32 converts string input to float32 slice.
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

// StringToFloat64 converts string input to float64 slice.
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

// ContainsString checks if all elements from needles list are in the haystack.
func ContainsString(haystack []string, needles ...string) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[string]struct{}, len(needles))
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

// ContainsAtLeastOneString checks if at least one element from needles list is in the haystack.
func ContainsAtLeastOneString(haystack []string, needles ...string) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if h == needles[0] {
				return true
			}
		}
		return false
	}

	checks := make(map[string]struct{}, len(needles))
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

// MatchesString checks if all elements from needles list are in the haystack in a case insensitive manner.
func MatchesString(haystack []string, needles ...string) bool {
	// Avoid allocations for a single check.
	if len(needles) == 1 {
		for _, h := range haystack {
			if strings.EqualFold(h, needles[0]) {
				return true
			}
		}
		return false
	}

	checks := make(map[string]struct{}, len(needles))
	for _, n := range needles {
		checks[strings.ToLower(n)] = struct{}{}
	}

	for _, h := range haystack {
		_, ok := checks[strings.ToLower(h)]
		if ok {
			delete(checks, h)

			if len(checks) == 0 {
				return true
			}
		}
	}

	return false
}

// TrimString removes empty string or whitespace only elements out of given string slice.
func TrimString(i []string) []string {
	var o []string
	for _, v := range i {
		if strings.TrimSpace(v) == "" {
			continue
		}

		o = append(o, v)
	}

	return o
}