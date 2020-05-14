package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestContainsUint64(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint64
		contains []uint64
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []uint64{1, 2, 3, 4, 5, 6},
			contains: []uint64{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []uint64{1, 2, 3, 4, 5, 6},
			contains: []uint64{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []uint64{1, 2, 3, 4, 5, 6},
			contains: []uint64{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []uint64{1, 2, 3, 4, 5, 6},
			contains: []uint64{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []uint64{1, 2, 3, 4, 5, 6},
			contains: []uint64{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsUint64(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsUint64_Single(b *testing.B) {
	haystack := []uint64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsUint64(haystack, 3)
	}
}

func BenchmarkContainsUint64_Multiple(b *testing.B) {
	haystack := []uint64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsUint64(haystack, 2, 3)
	}
}
