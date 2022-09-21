package netx

import (
	"encoding/binary"
	"net"
)

type IP net.IP

func (x *IP) ToIP() net.IP {
	return net.IP(*x)
}

func (x *IP) Uint32() uint32 {
	return binary.BigEndian.Uint32(x.to4())
}

func (x *IP) String() string {
	return net.IP(*x).String()
}

func (x *IP) Set(s string) error {
	if s == "" {
		return nil
	}

	*x = IP(net.ParseIP(s))

	return nil
}

func (x *IP) Type() string {
	return "IP"
}

func (x *IP) to4() net.IP {
	ip4 := net.IP(*x).To4()
	if len(ip4) != net.IPv4len {
		ip4 = net.ParseIP("0.0.0.0")
	}
	return ip4
}
