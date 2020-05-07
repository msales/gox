package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestContainsString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		contains []string
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a"},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a", "b", "c"},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"e", "f", "x", "y"},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"x", "y"},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsString(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		contains []int
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{7, 9},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsInt(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
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
