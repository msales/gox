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
			contains: []string{"a", "B", "C"},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A", "B"},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A"},
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

func TestMatchesString(t *testing.T) {
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
			name:     "found 1 element insensitive",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A"},
			want:     true,
		},
		{
			name:     "found 3 elements insensitive",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A", "B", "C"},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a", "x", "y"},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"x", "y"},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"x"},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatchesString(tt.slice, tt.contains...)
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
		{
			name:     "found no element",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{7},
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

func TestContainsUint(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		contains []uint
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []uint{1, 2, 3, 4, 5, 6},
			contains: []uint{3},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []uint{1, 2, 3, 4, 5, 6},
			contains: []uint{3, 5, 6},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []uint{1, 2, 3, 4, 5, 6},
			contains: []uint{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []uint{1, 2, 3, 4, 5, 6},
			contains: []uint{7, 9},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []uint{1, 2, 3, 4, 5, 6},
			contains: []uint{7},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsUint(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
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

func BenchmarkContainsString_Single(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsString(haystack, "c")
	}
}

func BenchmarkContainsString_Multiple(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsString(haystack, "b", "c")
	}
}

func BenchmarkMatchesString_Single(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MatchesString(haystack, "C")
	}
}

func BenchmarkMatchesString_Multiple(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MatchesString(haystack, "B", "C")
	}
}

func BenchmarkContainsInt_Single(b *testing.B) {
	haystack := []int{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsInt(haystack, 3)
	}
}

func BenchmarkContainsInt_Multiple(b *testing.B) {
	haystack := []int{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsInt(haystack, 2, 3)
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

func BenchmarkContainsUint_Single(b *testing.B) {
	haystack := []uint{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsUint(haystack, 3)
	}
}

func BenchmarkContainsUint_Multiple(b *testing.B) {
	haystack := []uint{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsUint(haystack, 2, 3)
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
