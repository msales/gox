package urlx_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/msales/gox/urlx"
)

func TestParseQuery(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    urlx.Values
		wantErr string // substring match; empty means no error expected
	}{
		{
			name: "empty query",
			in:   "",
			want: nil,
		},
		{
			name: "single pair",
			in:   "foo=bar",
			want: urlx.Values{{Key: "foo", Val: "bar"}},
		},
		{
			name: "multiple pairs preserve order",
			in:   "b=2&a=1&c=3",
			want: urlx.Values{
				{Key: "b", Val: "2"},
				{Key: "a", Val: "1"},
				{Key: "c", Val: "3"},
			},
		},
		{
			name: "duplicate keys kept in order",
			in:   "k=1&j=x&k=2",
			want: urlx.Values{
				{Key: "k", Val: "1"},
				{Key: "j", Val: "x"},
				{Key: "k", Val: "2"},
			},
		},
		{
			name: "key without equals gets empty value",
			in:   "foo",
			want: urlx.Values{{Key: "foo", Val: ""}},
		},
		{
			name: "empty segments are skipped",
			in:   "a=1&&b=2",
			want: urlx.Values{
				{Key: "a", Val: "1"},
				{Key: "b", Val: "2"},
			},
		},
		{
			name: "decodes plus and percent escapes",
			in:   "a+b=c%20d",
			want: urlx.Values{{Key: "a b", Val: "c d"}},
		},
		{
			name:    "semicolon in single pair returns error",
			in:      "a=1;b=2",
			want:    nil,
			wantErr: "semicolon",
		},
		{
			name: "semicolon skips bad segment but keeps valid ones",
			in:   "a=1&b=2;c&d=4",
			want: urlx.Values{
				{Key: "a", Val: "1"},
				{Key: "d", Val: "4"},
			},
			wantErr: "semicolon",
		},
		{
			name:    "invalid escape in key",
			in:      "a%ZZ=1",
			want:    nil,
			wantErr: "invalid URL escape",
		},
		{
			name:    "invalid escape in value",
			in:      "a=%ZZ",
			want:    nil,
			wantErr: "invalid URL escape",
		},
		{
			name: "first error reported but valid pairs still returned",
			in:   "a=%ZZ&b=2",
			want: urlx.Values{
				{Key: "b", Val: "2"},
			},
			wantErr: "invalid URL escape",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := urlx.ParseQuery(tt.in)

			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error = %v, want substring %q", err, tt.wantErr)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestParseQuery_EncodeRoundTrip(t *testing.T) {
	in := "b=2&a=1&k=1&k=2&empty="
	v, err := urlx.ParseQuery(in)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got := v.Encode(); got != in {
		t.Errorf("round-trip: got %q, want %q", got, in)
	}
}
