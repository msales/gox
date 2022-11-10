package gofp_test

import (
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestAny(t *testing.T) {
	tests := []struct {
		name       string
		containsFn func() bool
		want       bool
	}{
		{
			name: "found simple",
			containsFn: func() bool {
				slice := []int{1, 2, 3}
				return Any(slice, func(i int) bool {
					return i == 1
				})
			},
			want: true,
		},
		{
			name: "not found simple",
			containsFn: func() bool {
				slice := []int{1, 2, 3}
				return Any(slice, func(i int) bool {
					return i == 4
				})
			},
			want: false,
		},
		{
			name: "found struct",
			containsFn: func() bool {
				slice := []testStruct{
					{"1", 1},
					{"2", 2},
					{"3", 3},
				}
				return Any(slice, func(ts testStruct) bool {
					return ts.name == "1"
				})
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.containsFn()
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkAny_Simple(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Any(haystack, func(i int32) bool {
			return i == 3
		})
	}
}

func BenchmarkAny_Struct(b *testing.B) {
	haystack := []testStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"6", 5},
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Any(haystack, func(ts testStruct) bool {
			return ts.name == "4"
		})
	}
}

type testStruct struct {
	name  string
	value int
}
