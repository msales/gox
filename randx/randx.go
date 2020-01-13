package randx

import "math/rand"

// Happens returns true with a chance determined by a probability p.
// Probability should remain within 0.0 and 1.0 range.
// Probability lower or equal to 0.0 will always return false.
// Probability higher or equal to 1.0 will always return true.
func Happens(p float32) bool {
	if p <= 0 {
		return false
	}
	if p >= 1 {
		return true
	}
	return rand.Float32() < p
}

// WhichHappens returns index of a probability which occurred from a provided list. Probabilities are exclusive.
// Sum of probabilities should not exceed 1.0 and each probability should remain within 0.0 and 1.0 range.
// If one or more probabilities are outside of allowed range received results will not be valid.
func WhichHappens(p ...float32) int {
	chance := rand.Float32()
	var sum float32
	for i, probability := range p {
		sum += probability
		if chance < sum {
			return i
		}
	}

	return -1
}