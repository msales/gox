package gofp_test

import (
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestContainsAll(t *testing.T) {
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
		{
			name:     "found all in equal slices",
			slice:    []int64{1, 2, 3},
			contains: []int64{1, 2, 3},
			want:     true,
		},
		{
			name:     "found some with needles containing all slice elements plus some",
			slice:    []int64{1, 2, 3},
			contains: []int64{1, 2, 3, 4},
			want:     false,
		},
		{
			name:     "found no element in empty slicee",
			slice:    []int64{},
			contains: []int64{1, 2, 3, 4},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsAll(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAll_Single(b *testing.B) {
	haystack := []int64{1, 2, 3, 4, 5, 6}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAll(haystack, 3)
	}
}

func BenchmarkContainsAll_Multiple(b *testing.B) {
	haystack := []int64{1, 2, 3, 4, 5, 6}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAll(haystack, 2, 3)
	}
}

func TestContainsAtLeastOne(t *testing.T) {
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
			want:     true,
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
			got := ContainsAtLeastOne(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAtLeastOne_Single(b *testing.B) {
	haystack := []int64{1, 2, 3, 4, 5, 6}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOne(haystack, 3)
	}
}

func BenchmarkContainsAtLeastOne_Multiple(b *testing.B) {
	haystack := []int64{1, 2, 3, 4, 5, 6}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOne(haystack, 2, 3)
	}
}
