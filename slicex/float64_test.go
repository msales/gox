package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestFloat64ToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []float64{1.3333, 2.666, 3.1},
			want:  []interface{}{float64(1.3333), float64(2.666), float64(3.1)},
		},
		{
			name:  "empty input",
			input: []float64{},
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
			got := Float64ToInterface(tt.input)

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

func TestFloat64ToString(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		precision int
		want      []string
	}{
		{
			name:      "correct input",
			input:     []float64{1.3333, 2.666, 3.1},
			precision: 2,
			want:      []string{"1.33", "2.67", "3.10"},
		},
		{
			name:      "empty input",
			input:     []float64{},
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
			got := Float64ToString(tt.input, tt.precision)

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

func TestFloat64ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		want  []int
	}{
		{
			name:  "correct input",
			input: []float64{1.3333, 2.666, 3.1},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []float64{},
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
			got := Float64ToInt(tt.input)

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

func TestFloat64ToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		want  []int32
	}{
		{
			name:  "correct input",
			input: []float64{1.3333, 2.666, 3.1},
			want:  []int32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []float64{},
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
			got := Float64ToInt32(tt.input)

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

func TestFloat64ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		want  []int64
	}{
		{
			name:  "correct input",
			input: []float64{1.3333, 2.666, 3.1},
			want:  []int64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []float64{},
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
			got := Float64ToInt64(tt.input)

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

func TestFloat64ToFloat32(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		want  []float32
	}{
		{
			name:  "correct input",
			input: []float64{1.3333, 2.666, 3.1},
			want:  []float32{float32(float64(1.3333)), float32(float64(2.666)), float32(float64(3.1))},
		},
		{
			name:  "empty input",
			input: []float64{},
			want:  []float32{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []float32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64ToFloat32(tt.input)

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

func TestContainsFloat64(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		contains []float64
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []float64{1, 2, 3, 4, 5, 6},
			contains: []float64{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []float64{1, 2, 3, 4, 5, 6},
			contains: []float64{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []float64{1, 2, 3, 4, 5, 6},
			contains: []float64{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []float64{1, 2, 3, 4, 5, 6},
			contains: []float64{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []float64{1, 2, 3, 4, 5, 6},
			contains: []float64{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsFloat64(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsFloat64_Single(b *testing.B) {
	haystack := []float64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsFloat64(haystack, 3)
	}
}

func BenchmarkContainsFloat64_Multiple(b *testing.B) {
	haystack := []float64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsFloat64(haystack, 2, 3)
	}
}
