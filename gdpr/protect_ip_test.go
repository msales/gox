package gdpr_test

import (
	"net"
	"strconv"
	"testing"

	. "github.com/msales/gox/gdpr"
	"github.com/msales/gox/netx"
)

func Test_ProtectIP_netIP(t *testing.T) {
	tests := []struct {
		name string
		ip   net.IP
		want string
	}{
		{
			name: "nil",
			ip:   nil,
			want: "",
		},
		{
			name: "empty",
			ip:   net.IP{},
			want: "",
		},
		{
			name: "net.IP type",
			ip:   net.IPv4(42, 21, 37, 213),
			want: "42.21.37.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtectIP(tt.ip); got != tt.want {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_ProtectIP_netxIP(t *testing.T) {
	tests := []struct {
		name string
		ip   netx.IP
		want string
	}{
		{
			name: "nil",
			ip:   nil,
			want: "",
		},
		{
			name: "empty",
			ip:   netx.IP{},
			want: "",
		},
		{
			name: "net.IP type",
			ip:   netx.IP(net.IPv4(42, 21, 37, 213)),
			want: "42.21.37.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtectIP(tt.ip); got != tt.want {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_ProtectIP_String(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want string
	}{
		{
			name: "empty string",
			ip:   "",
			want: "",
		},
		{
			name: "zero ip",
			ip:   "0.0.0.0",
			want: "0.0.0.0",
		},
		{
			name: "raw ip v4",
			ip:   "42.21.37.213",
			want: "42.21.37.0",
		},
		{
			name: "local ip v6",
			ip:   "::1",
			want: "::1",
		},
		{
			name: "raw ip v6",
			ip:   "2001:0000:130F:0000:0000:09C0:876A:130B",
			want: "2001:0000:130f:0000:0000:09c0:0000:0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtectIP(tt.ip); got != tt.want {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_ProtectIP_Uint(t *testing.T) {
	tests := []struct {
		name string
		ip   uint32
		want string
	}{
		{
			name: "0 value",
			ip:   0,
			want: "",
		},
		{
			name: "uint ip",
			ip:   ipToUint(net.IPv4(42, 21, 37, 213)),
			want: "42.21.37.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtectIP(tt.ip); got != tt.want {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func ipToUint(ip net.IP) uint32 {
	xIP := netx.IP(ip)
	return xIP.Uint32()
}

func BenchmarkProtectIP_Uint(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectIP(uint32(i))
	}
}

func BenchmarkProtectIP_String(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		v := strconv.Itoa(i)
		ProtectIP(v + "." + v + "." + v + "." + v)
	}
}

func BenchmarkProtectIP_netIP(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectIP(net.IPv4(byte(i), byte(i), byte(i), byte(i)))
	}
}

func BenchmarkProtectIP_netxIP(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectIP(netx.IP(net.IPv4(byte(i), byte(i), byte(i), byte(i))))
	}
}
