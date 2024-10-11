package gdpr_test

import (
	"strconv"
	"testing"

	. "github.com/msales/gox/gdpr"
)

func Test_ProtectDeviceID(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "empty value",
			value: "",
			want:  "",
		},
		{
			name:  "correct value to protect",
			value: "some_value",
			want:  "some_val**",
		},
		{
			name:  "only numbers",
			value: "12345",
			want:  "123**",
		},
		{
			name:  "string with 2 chars",
			value: "11",
			want:  "**",
		},
		{
			name:  "string with less than 2 chars",
			value: "1",
			want:  "*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtectDeviceID(tt.value); got != tt.want {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkProtectDeviceID(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectDeviceID(strconv.Itoa(i))
	}
}
