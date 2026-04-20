package urlx_test

import (
	"testing"

	"github.com/msales/gox/urlx"
)

func TestValues_Encode(t *testing.T) {
	tests := []struct {
		name string
		in   urlx.Values
		want string
	}{
		{
			name: "nil slice",
			in:   nil,
			want: "",
		},
		{
			name: "empty slice",
			in:   urlx.Values{},
			want: "",
		},
		{
			name: "single pair",
			in:   urlx.Values{{Key: "foo", Val: "bar"}},
			want: "foo=bar",
		},
		{
			name: "preserves insertion order",
			in: urlx.Values{
				{Key: "b", Val: "2"},
				{Key: "a", Val: "1"},
				{Key: "c", Val: "3"},
			},
			want: "b=2&a=1&c=3",
		},
		{
			name: "duplicate keys kept in order",
			in: urlx.Values{
				{Key: "k", Val: "1"},
				{Key: "j", Val: "x"},
				{Key: "k", Val: "2"},
			},
			want: "k=1&j=x&k=2",
		},
		{
			name: "escapes space in key and value",
			in:   urlx.Values{{Key: "a b", Val: "c d"}},
			want: "a+b=c+d",
		},
		{
			name: "escapes reserved characters",
			in:   urlx.Values{{Key: "&", Val: "="}},
			want: "%26=%3D",
		},
		{
			name: "empty value keeps equals sign",
			in:   urlx.Values{{Key: "k", Val: ""}},
			want: "k=",
		},
		{
			name: "empty key",
			in:   urlx.Values{{Key: "", Val: "v"}},
			want: "=v",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.Encode()
			if got != tt.want {
				t.Errorf("Encode() = %q, want %q", got, tt.want)
			}
		})
	}
}
