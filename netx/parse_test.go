package netx

import "testing"

func TestRawIPToUint(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want uint32
	}{
		{
			name: "incorrect IP - 123",
			raw:  "123",
			want: 0,
		},
		{
			name: "incorrect IP - test",
			raw:  "test",
			want: 0,
		},
		{
			name: "correct IP",
			raw:  "127.0.0.1",
			want: uint32(0b01111111000000000000000000000001),
		},
		{
			name: "correct IPv6 - not supported",
			raw:  "2001:db8:3333:4444:5555:6666:7777:8888",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RawIPToUint(tt.raw); got != tt.want {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestUintToIP(t *testing.T) {
	tests := []struct {
		name   string
		uintIP uint32
		want   string
	}{
		{
			name:   "incorrect IP",
			uintIP: 0,
			want:   "0.0.0.0",
		},
		{
			name:   "correct IP",
			uintIP: uint32(0b01111111000000000000000000000001),
			want:   "127.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintToIP(tt.uintIP); got.String() != tt.want {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}
