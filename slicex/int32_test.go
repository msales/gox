package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestInt32ToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []int32{1, 2, 3},
			want:  []interface{}{int32(1), int32(2), int32(3)},
		},
		{
			name:  "empty input",
			input: []int32{},
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
			got := Int32ToInterface(tt.input)

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

func TestInt32ToString(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  []string
	}{
		{
			name:  "correct input",
			input: []int32{1, 2, 3},
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "empty input",
			input: []int32{},
			want:  []string{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32ToString(tt.input)

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

func TestInt32ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  []int
	}{
		{
			name:  "correct input",
			input: []int32{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int32{},
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
			got := Int32ToInt(tt.input)

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

func TestInt32ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  []int64
	}{
		{
			name:  "correct input",
			input: []int32{1, 2, 3},
			want:  []int64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int32{},
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
			got := Int32ToInt64(tt.input)

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

func TestInt32ToFloat32(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  []float32
	}{
		{
			name:  "correct input",
			input: []int32{1, 2, 3},
			want:  []float32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int32{},
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
			got := Int32ToFloat32(tt.input)

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

func TestInt32ToFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  []float64
	}{
		{
			name:  "correct input",
			input: []int32{1, 2, 3},
			want:  []float64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int32{},
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
			got := Int32ToFloat64(tt.input)

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

func TestContainsInt32(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int32
		contains []int32
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsInt32(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsInt32_Single(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsInt32(haystack, 3)
	}
}

func BenchmarkContainsInt32_Multiple(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsInt32(haystack, 2, 3)
	}
}

func TestContainsAtLeastOneInt32(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int32
		contains []int32
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{3, 5, 6, 7},
			want:     true,
		},
		{
			name:     "found no elements",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []int32{1, 2, 3, 4, 5, 6},
			contains: []int32{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsAtLeastOneInt32(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAtLeastOneInt32_Single(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneInt32(haystack, 3)
	}
}

func BenchmarkContainsAtLeastOneInt32_Multiple(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneInt32(haystack, 2, 3)
	}
}
