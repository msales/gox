package mathx

import (
	"math"
)

// Round rounds passed value to number of decimals.
func Round(f, decimals float64) float64 {
	roundTo := math.Pow(10, decimals)

	return math.Round(roundTo*f) / roundTo
}
