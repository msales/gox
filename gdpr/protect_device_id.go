package gdpr

const (
	hiddenRune = '*'
)

// ProtectDeviceID hides last two character from passed device id and returns string with protected value
func ProtectDeviceID(val string) string {
	if val == "" {
		return val
	}

	r := []rune(val)
	l := len(r)

	// If someone passes string with less than 2 characters, we don't protect it.
	if l < 2 {
		return val
	}

	r[l-1] = hiddenRune
	r[l-2] = hiddenRune

	return string(r)
}
