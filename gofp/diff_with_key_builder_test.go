package gofp_test

import (
	"reflect"
	"testing"

	"github.com/msales/gox/gofp"
	"github.com/stretchr/testify/assert"
)

func TestDiffLeftWithKeyBuilder(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		kb     gofp.KeyBuilder[int, int]
		want   []int
	}{
		{
			name:   "disjoint slices",
			slice1: []int{1, 2, 3},
			slice2: []int{4, 5, 6},
			kb:     func(i int) int { return i },
			want:   []int{1, 2, 3},
		},
		{
			name:   "overlapping slices",
			slice1: []int{1, 2, 3},
			slice2: []int{3, 4, 5},
			kb:     func(i int) int { return i },
			want:   []int{1, 2},
		},
		{
			name:   "identical slices",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2, 3},
			kb:     func(i int) int { return i },
			want:   []int{},
		},
		{
			name:   "repeated values",
			slice1: []int{1, 1, 1, 1, 2, 2, 2, 1},
			slice2: []int{3, 3, 3, 3, 3, 3, 3},
			kb:     func(i int) int { return i },
			want:   []int{1, 2},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gofp.DiffLeftWithKeyBuilder(tt.slice1, tt.slice2, tt.kb)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffLeftWithKeyBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffRightWithKeyBuilder(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		kb     gofp.KeyBuilder[int, int]
		want   []int
	}{
		{
			name:   "disjoint slices",
			slice1: []int{4, 5, 6},
			slice2: []int{1, 2, 3},
			kb:     func(i int) int { return i },
			want:   []int{1, 2, 3},
		},
		{
			name:   "overlapping slices",
			slice1: []int{3, 4, 5},
			slice2: []int{1, 2, 3},
			kb:     func(i int) int { return i },
			want:   []int{1, 2},
		},
		{
			name:   "identical slices",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2, 3},
			kb:     func(i int) int { return i },
			want:   []int{},
		},
		{
			name:   "repeated values",
			slice1: []int{3, 3, 3, 3, 3, 3, 3},
			slice2: []int{1, 1, 1, 1, 2, 2, 2, 1},
			kb:     func(i int) int { return i },
			want:   []int{1, 2},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gofp.DiffRightWithKeyBuilder(tt.slice1, tt.slice2, tt.kb)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffLeftWithKeyBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymDiffWithKeyBuilder(t *testing.T) {
	// define the test cases
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		kb     gofp.KeyBuilder[int, int]
		want   []int
	}{
		{
			name:   "disjoint slices",
			slice1: []int{1, 2, 3},
			slice2: []int{4, 5, 6},
			kb:     func(i int) int { return i },
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "overlapping slices",
			slice1: []int{1, 2, 3},
			slice2: []int{3, 4, 5},
			kb:     func(i int) int { return i },
			want:   []int{1, 2, 4, 5},
		},
		{
			name:   "identical slices",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2, 3},
			kb:     func(i int) int { return i },
			want:   []int{},
		},
		{
			name:   "repeated values",
			slice1: []int{1, 1, 1, 1, 2, 2, 2, 1},
			slice2: []int{3, 3, 3, 3, 3, 3, 3},
			kb:     func(i int) int { return i },
			want:   []int{1, 2, 3},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gofp.SymDiffWithKeyBuilder(tt.slice1, tt.slice2, tt.kb)
			for _, v := range tt.want {
				assert.Contains(t, got, v, "Got %+v, want %+v")
			}
		})
	}
}
