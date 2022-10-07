package gofp_test

import (
	"sort"
	"testing"

	"github.com/msales/gox/gofp"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		name   string
		slices [][]int
		want   []int
	}{
		{
			name: "One common element",
			slices: [][]int{
				{1, 3, 5},
				{2, 5, 7},
			},
			want: []int{1, 2, 3, 5, 7},
		},
		{
			name: "Multiple common elements",
			slices: [][]int{
				{1, 3, 5, 7},
				{3, 5, 7, 9},
			},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "More than two slices",
			slices: [][]int{
				{1, 3, 5, 6},
				{3, 5, 7, 9},
				{3, 5, 4, 10, 7},
				{2, 5, 4, 10, 11},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11},
		},
		{
			name: "No common elements",
			slices: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "All elements in common",
			slices: [][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gofp.Union(tt.slices...)

			sort.Ints(tt.want)
			sort.Ints(result)

			if len(tt.want) != len(result) {
				t.Errorf("Got %+v, want %+v", result, tt.want)
				return
			}

			for idx, value := range tt.want {
				if value != result[idx] {
					t.Errorf("Got %+v, want %+v", result, tt.want)
					break
				}
			}
		})
	}
}

func BenchmarkUnion_TwoSlices(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 6, 7, 8, 9}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.Union(slice1, slice2)
	}
}

func BenchmarkUnion_ManySlices(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 6, 7, 8, 9}
	slice3 := []int{9, 10, 11, 12, 13}
	slice4 := []int{13, 14, 15, 16, 17}
	slice5 := []int{18, 19, 20, 21, 22}
	slice6 := []int{23, 24, 25, 26, 27}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.Union(slice1, slice2, slice3, slice4, slice5, slice6)
	}
}
