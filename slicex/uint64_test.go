package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestUint64ToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []uint64
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []uint64{1, 2, 3},
			want:  []interface{}{uint64(1), uint64(2), uint64(3)},
		},
		{
			name:  "empty input",
			input: []uint64{},
			want:  []interface{}{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint64ToInterface(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for k, v := range tt.want {
				if v != got[k] {
					t.Errorf("Got %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

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

func TestContainsAtLeastOneUint64(t *testing.T) {
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
			want:     true,
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
			got := ContainsAtLeastOneUint64(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAtLeastOneUint64_Single(b *testing.B) {
	haystack := []uint64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneUint64(haystack, 3)
	}
}

func BenchmarkContainsAtLeastOneUint64_Multiple(b *testing.B) {
	haystack := []uint64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneUint64(haystack, 2, 3)
	}
}
