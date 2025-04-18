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
			name:  "unknown value - lower case",
			value: "unknown",
			want:  "unknown",
		},
		{
			name:  "unknown value - upper case",
			value: "UNKNOWN",
			want:  "UNKNOWN",
		},
		{
			name:  "unknown value - mixed case",
			value: "unKNowN",
			want:  "unKNowN",
		},
		{
			name:  "non available value - upper case",
			value: "N/A",
			want:  "N/A",
		},
		{
			name:  "non available value - lower case",
			value: "n/a",
			want:  "n/a",
		},
		{
			name:  "non available value - mixed case",
			value: "n/A",
			want:  "n/A",
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

func BenchmarkProtectDeviceID_Empty(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectDeviceID("")
	}
}

func BenchmarkProtectDeviceID_NA_Mixed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectDeviceID("N/a")
	}
}

func BenchmarkProtectDeviceID_NA_Lower(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectDeviceID("n/a")
	}
}

func BenchmarkProtectDeviceID_Unknown_Mixed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectDeviceID("Unknown")
	}
}

func BenchmarkProtectDeviceID_Unknown_Lower(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ProtectDeviceID("unknown")
	}
}
