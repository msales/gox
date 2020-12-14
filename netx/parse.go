package netx

import (
	"encoding/binary"
	"net"
)

// RawIPToUint convert raw IP e.g. 127.0.0.1 to uint32 number.
func RawIPToUint(raw string) uint32 {
	ip := net.ParseIP(raw)
	if ip == nil {
		return 0
	}

	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// UintToIP converts uint32 number to IP e.g. 127.0.0.1.
func UintToIP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
