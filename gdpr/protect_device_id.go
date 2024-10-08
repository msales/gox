package gdpr

import "strings"

const (
	ipSeparator         = "."
	gdprStringHideValue = "*"
	gdprIPHideValue     = "0"
	unknownIPValue      = "unknown"
	emptyIP             = "0"
)

// ProtectDeviceID hides last two character from passed device id and returns string with protected value
func ProtectDeviceID(deviceIDValue string) string {
	if deviceIDValue == "" {
		return deviceIDValue
	}

	splitted := strings.Split(deviceIDValue, "")
	l := len(splitted)
	splitted[l-1] = gdprStringHideValue
	splitted[l-2] = gdprStringHideValue

	return strings.Join(splitted, "")
}
