package gdpr

import (
	"strings"

	"github.com/msales/gox/netx"
)

// ProtectIPV4 hides last octet from passed uint32 IPV4 value and returns string with protected value
func ProtectIPV4(ipValue uint32) string {
	if ipValue == 0 {
		return netx.UintToIP(ipValue).String()
	}

	ip := netx.UintToIP(ipValue).String()

	return ProtectRawIP(ip)
}

// ProtectRawIP hides last octet from ip string value and returns string with protected value
func ProtectRawIP(ipValue string) string {
	if ipValue == emptyIP || ipValue == "" || ipValue == unknownIPValue {
		return ipValue
	}

	splitted := strings.Split(ipValue, ".")
	l := len(splitted)
	splitted[l-1] = gdprIPHideValue

	return strings.Join(splitted, ipSeparator)
}
