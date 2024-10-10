package gdpr

import (
	"encoding/binary"
	"net"
	"net/netip"

	"github.com/msales/gox/netx"
)

// IP contains constraint for accepted types of IPs for protection.
type IP interface {
	net.IP | netx.IP | uint32 | string
}

// ProtectIP hides last octet from passed uint32 IPV4 value and returns string with protected value
func ProtectIP[T IP](ip T) string {
	switch cIP := any(ip).(type) {
	// Every other type should end up in this case.
	case net.IP:
		// if ip is somehow empty or nil, return empty string
		if len(cIP) == 0 {
			return ""
		}

		// if IP is v6, we don't do anything with it.
		if cIP.To4() == nil {

			// if IPv6 is local we don't do anything
			if cIP.String() == "::1" {
				return cIP.String()
			}

			cIP[12] = 0
			cIP[13] = 0
			cIP[14] = 0
			cIP[15] = 0

			addr, err := netip.ParseAddr(cIP.String())
			// if there is an error, we don't do anything, we return empty value because something is wrong with passed ip
			if err != nil {
				return ""
			}

			return addr.StringExpanded()
		}

		// change last octet of IP v4 to 0 and guard
		cIP[15] = 0
		return cIP.String()

	// netx.IP is just wrapper over net.IP type, so it's easy to protect against it.
	case netx.IP:
		return ProtectIP(cIP.ToIP())

	// uint32 is number representation of IP used by gox/netx package. 0 is the only edge value to check
	case uint32:
		if cIP == 0 {
			return ""
		}

		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b[:], cIP)
		return ProtectIP(net.IPv4(b[0], b[1], b[2], b[3]))

	// string is raw representation of IP, downside is that we have to parse it first.
	case string:
		if cIP == "" {
			return ""
		}

		parsed := net.ParseIP(cIP)
		return ProtectIP(parsed)
	}

	return ""
}
