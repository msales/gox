package mathx

import (
	"math"
)

func Round(f, decimals float64) float64 {
	roundTo := math.Pow(10, decimals)

	return math.Round(roundTo*f) / roundTo
}
