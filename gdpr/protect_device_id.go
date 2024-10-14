package gdpr

import "strings"

const (
	hiddenRune        = '*'
	unknownValue      = "unknown"
	nonAvailableValue = "n/a"
)

// ProtectDeviceID hides last two character from passed device id and returns string with protected value
func ProtectDeviceID(val string) string {
	lowered := strings.ToLower(val)

	if val == "" || lowered == unknownValue || lowered == nonAvailableValue {
		return val
	}

	r := []rune(val)
	l := len(r)

	// If someone passes string with less than 2 characters, protect everything.
	c := 2
	if l < 2 {
		c = l
	}

	for i := 1; i <= c; i++ {
		r[l-i] = hiddenRune
	}

	return string(r)
}
