package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestUint32ToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []uint32
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []uint32{1, 2, 3},
			want:  []interface{}{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []uint32{},
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
			got := Uint32ToInterface(tt.input)

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

func TestContainsUint32(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		contains []uint32
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []uint32{1, 2, 3, 4, 5, 6},
			contains: []uint32{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []uint32{1, 2, 3, 4, 5, 6},
			contains: []uint32{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []uint32{1, 2, 3, 4, 5, 6},
			contains: []uint32{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []uint32{1, 2, 3, 4, 5, 6},
			contains: []uint32{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []uint32{1, 2, 3, 4, 5, 6},
			contains: []uint32{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsUint32(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsUint32_Single(b *testing.B) {
	haystack := []uint32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsUint32(haystack, 3)
	}
}

func BenchmarkContainsUint32_Multiple(b *testing.B) {
	haystack := []uint32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsUint32(haystack, 2, 3)
	}
}
