package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestInt64ToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []int64{1, 2, 3},
			want:  []interface{}{int64(1), int64(2), int64(3)},
		},
		{
			name:  "empty input",
			input: []int64{},
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
			got := Int64ToInterface(tt.input)

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

func TestInt64ToString(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []string
	}{
		{
			name:  "correct input",
			input: []int64{1, 2, 3},
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "empty input",
			input: []int64{},
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
			got := Int64ToString(tt.input)

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

func TestInt64ToInt(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int
	}{
		{
			name:  "correct input",
			input: []int64{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int64{},
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
			got := Int64ToInt(tt.input)

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

func TestInt64ToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int32
	}{
		{
			name:  "correct input",
			input: []int64{1, 2, 3},
			want:  []int32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int64{},
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
			got := Int64ToInt32(tt.input)

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

func TestInt64ToFloat32(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []float32
	}{
		{
			name:  "correct input",
			input: []int64{1, 2, 3},
			want:  []float32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int64{},
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
			got := Int64ToFloat32(tt.input)

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

func TestInt64ToFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []float64
	}{
		{
			name:  "correct input",
			input: []int64{1, 2, 3},
			want:  []float64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int64{},
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
			got := Int64ToFloat64(tt.input)

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

func TestContainsInt64(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int64
		contains []int64
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []int64{1, 2, 3, 4, 5, 6},
			contains: []int64{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []int64{1, 2, 3, 4, 5, 6},
			contains: []int64{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []int64{1, 2, 3, 4, 5, 6},
			contains: []int64{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []int64{1, 2, 3, 4, 5, 6},
			contains: []int64{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []int64{1, 2, 3, 4, 5, 6},
			contains: []int64{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsInt64(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsInt64_Single(b *testing.B) {
	haystack := []int64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsInt64(haystack, 3)
	}
}

func BenchmarkContainsInt64_Multiple(b *testing.B) {
	haystack := []int64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsInt64(haystack, 2, 3)
	}
}
