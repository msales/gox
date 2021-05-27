package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name       string
		containsFn func() bool
		want       bool
	}{
		{
			name: "found simple",
			containsFn: func() bool {
				slice := []int{1, 2, 3}
				return Contains(slice, func(i int) bool {
					return slice[i] == 1
				})
			},
			want: true,
		},
		{
			name: "not found simple",
			containsFn: func() bool {
				slice := []int{1, 2, 3}
				return Contains(slice, func(i int) bool {
					return slice[i] == 4
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
				return Contains(slice, func(i int) bool {
					return slice[i].name == "1"
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

func TestContains_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Contains should panic when not slice is passed")
		}
	}()

	Contains(1, func(i int) bool {
		return true
	})
}

func BenchmarkContains_Simple(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Contains(haystack, func(i int) bool {
			return haystack[i] == 3
		})
	}
}

func BenchmarkContains_Struct(b *testing.B) {
	haystack := []testStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"6", 5},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Contains(haystack, func(i int) bool {
			return haystack[i].name == "4"
		})
	}
}

type testStruct struct {
	name  string
	value int
}
