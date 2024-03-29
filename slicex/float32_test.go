package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestFloat32ToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []float32
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []float32{1.3333, 2.666, 3.1},
			want:  []interface{}{float32(1.3333), float32(2.666), float32(3.1)},
		},
		{
			name:  "empty input",
			input: []float32{},
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
			got := Float32ToInterface(tt.input)

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

func TestFloat32ToString(t *testing.T) {
	tests := []struct {
		name      string
		input     []float32
		precision int
		want      []string
	}{
		{
			name:      "correct input",
			input:     []float32{1.3333, 2.666, 3.1},
			precision: 2,
			want:      []string{"1.33", "2.67", "3.10"},
		},
		{
			name:      "empty input",
			input:     []float32{},
			precision: 2,
			want:      []string{},
		},
		{
			name:      "nil input",
			input:     nil,
			precision: 2,
			want:      []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32ToString(tt.input, tt.precision)

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

func TestFloat32ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input []float32
		want  []int
	}{
		{
			name:  "correct input",
			input: []float32{1.3333, 2.666, 3.1},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []float32{},
			want:  []int{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32ToInt(tt.input)

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

func TestFloat32ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input []float32
		want  []int32
	}{
		{
			name:  "correct input",
			input: []float32{1.3333, 2.666, 3.1},
			want:  []int32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []float32{},
			want:  []int32{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []int32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32ToInt32(tt.input)

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

func TestFloat32ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input []float32
		want  []int64
	}{
		{
			name:  "correct input",
			input: []float32{1.3333, 2.666, 3.1},
			want:  []int64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []float32{},
			want:  []int64{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32ToInt64(tt.input)

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

func TestFloat32ToFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input []float32
		want  []float64
	}{
		{
			name:  "correct input",
			input: []float32{1.3333, 2.666, 3.1},
			want:  []float64{float64(float32(1.3333)), float64(float32(2.666)), float64(float32(3.1))},
		},
		{
			name:  "empty input",
			input: []float32{},
			want:  []float64{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32ToFloat64(tt.input)

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

func TestContainsFloat32(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float32
		contains []float32
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsFloat32(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestContainsAtLeastOneFloat32(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float32
		contains []float32
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{3, 5, 6, 7},
			want:     true,
		},
		{
			name:     "found no elements",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []float32{1, 2, 3, 4, 5, 6},
			contains: []float32{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsAtLeastOneFloat32(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsFloat32_Single(b *testing.B) {
	haystack := []float32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsFloat32(haystack, 3)
	}
}

func BenchmarkContainsFloat32_Multiple(b *testing.B) {
	haystack := []float32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsFloat32(haystack, 2, 3)
	}
}

func BenchmarkContainsAtLeastOneFloat32_Single(b *testing.B) {
	haystack := []float32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneFloat32(haystack, 3)
	}
}

func BenchmarkContainsAtLeastOneFloat32_Multiple(b *testing.B) {
	haystack := []float32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneFloat32(haystack, 2, 3)
	}
}
