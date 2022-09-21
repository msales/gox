package netx_test

import (
	"net"
	"reflect"
	"testing"

	"github.com/msales/gox/netx"
)

func TestIP_Set(t *testing.T) {
	tests := []struct {
		env  string
		want net.IP
	}{
		{
			env:  "",
			want: net.IP{},
		},
		{
			env:  "127.0.0.1",
			want: net.ParseIP("127.0.0.1"),
		},
		{
			env:  "2001:db8:3333:4444:5555:6666:7777:8888",
			want: net.ParseIP("2001:db8:3333:4444:5555:6666:7777:8888"),
		},
		{
			env:  "foobar",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			got := netx.IP{}
			err := got.Set(tt.env)

			if !reflect.DeepEqual(net.IP(got), tt.want) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
			if err != nil {
				t.Errorf("expecting no error, got %+v", err)
			}
		})
	}
}
