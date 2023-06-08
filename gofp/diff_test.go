package gofp

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymDiff(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{
			name:   "diff left slice",
			slice1: []int{1, 2, 3},
			slice2: []int{3},
			want:   []int{1, 2},
		},
		{
			name:   "diff right slice",
			slice1: []int{3},
			slice2: []int{1, 2, 3},
			want:   []int{1, 2},
		},
		{
			name:   "diff both slices",
			slice1: []int{3, 4, 5},
			slice2: []int{1, 2, 3},
			want:   []int{4, 5, 1, 2},
		},
		{
			name:   "diff both slices - same values",
			slice1: []int{1, 2, 3, 4, 5, 6},
			slice2: []int{1, 2, 3, 4, 5, 6},
			want:   []int{},
		},
		{
			name:   "diff empty",
			slice1: []int{},
			slice2: []int{},
			want:   []int{},
		},
		{
			name:   "diff both - pessimistic",
			slice1: []int{1, 2, 3, 4, 5, 6},
			slice2: []int{7, 8, 9, 10, 11, 12},
			want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name:   "repeated values",
			slice1: []int{1, 1, 1, 1, 2, 2, 2, 1},
			slice2: []int{3, 3, 3, 3, 3, 3, 3},
			want:   []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SymDiff(tt.slice1, tt.slice2)
			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for _, v := range tt.want {
				assert.Contains(t, got, v, "Got %+v, want %+v")
			}
		})
	}
}

func BenchmarkSymDiff_Both(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{5, 6, 7, 8, 9, 10}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SymDiff(slice1, slice2)
	}
}

func BenchmarkSymDiff_Both_Pessimistic(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SymDiff(slice1, slice2)
	}
}

func BenchmarkSymDiff_Both_MoreValues_Random(b *testing.B) {
	var slice1, slice2 []int
	for i := 0; i < 300; i++ {
		slice1 = append(slice1, i)
	}

	for i := 0; i < 50; i++ {
		slice2 = append(slice2, i)
	}
	rand.Shuffle(len(slice1), func(i, j int) { slice1[i], slice1[j] = slice1[j], slice1[i] })
	rand.Shuffle(len(slice2), func(i, j int) { slice2[i], slice2[j] = slice2[j], slice2[i] })

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SymDiff(slice1, slice2)
	}
}

func BenchmarkSymDiff_Both_Pessimistic_RightShorter(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	slice2 := []int{17, 18, 19, 20}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SymDiff(slice1, slice2)
	}
}

func TestDiffLeft(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{
			name:   "diff left slice",
			slice1: []int{1, 2, 3},
			slice2: []int{3},
			want:   []int{1, 2},
		},
		{
			name:   "diff right slice",
			slice1: []int{3},
			slice2: []int{1, 2, 3},
			want:   []int{},
		},
		{
			name:   "diff both",
			slice1: []int{3, 4, 5},
			slice2: []int{1, 2, 3},
			want:   []int{4, 5},
		},
		{
			name:   "diff both slices - same values",
			slice1: []int{1, 2, 3, 4, 5, 6},
			slice2: []int{1, 2, 3, 4, 5, 6},
			want:   []int{},
		},
		{
			name:   "diff empty",
			slice1: []int{},
			slice2: []int{},
			want:   []int{},
		},
		{
			name:   "diff both - pessimistic",
			slice1: []int{1, 2, 3, 4, 5, 6},
			slice2: []int{7, 8, 9, 10, 11, 12},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "diff both - pessimistic",
			slice1: []int{1, 2, 3, 4, 5, 6},
			slice2: []int{7, 8, 9, 10, 11, 12},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "repeated values",
			slice1: []int{1, 1, 1, 1, 2, 2, 2, 1},
			slice2: []int{3, 3, 3, 3, 3, 3, 3},
			want:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DiffLeft(tt.slice1, tt.slice2)
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

func BenchmarkDiffLeft(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{5, 6, 7, 8}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DiffLeft(slice1, slice2)
	}
}

func BenchmarkDiffLeft_MoreValues_Random(b *testing.B) {
	var slice1, slice2 []int
	for i := 0; i < 300; i++ {
		slice1 = append(slice1, i)
	}

	for i := 0; i < 50; i++ {
		slice2 = append(slice2, i)
	}
	rand.Shuffle(len(slice1), func(i, j int) { slice1[i], slice1[j] = slice1[j], slice1[i] })
	rand.Shuffle(len(slice2), func(i, j int) { slice2[i], slice2[j] = slice2[j], slice2[i] })

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DiffLeft(slice1, slice2)
	}
}

func TestDiffRight(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{
			name:   "diff left slice",
			slice1: []int{1, 2, 3},
			slice2: []int{3},
			want:   []int{},
		},
		{
			name:   "diff right slice",
			slice1: []int{3},
			slice2: []int{1, 2, 3},
			want:   []int{1, 2},
		},
		{
			name:   "diff both",
			slice1: []int{3, 4, 5},
			slice2: []int{1, 2, 3},
			want:   []int{1, 2},
		},
		{
			name:   "diff both slices - same values",
			slice1: []int{1, 2, 3, 4, 5, 6},
			slice2: []int{1, 2, 3, 4, 5, 6},
			want:   []int{},
		},
		{
			name:   "diff empty",
			slice1: []int{},
			slice2: []int{},
			want:   []int{},
		},
		{
			name:   "diff both - pessimistic",
			slice1: []int{1, 2, 3, 4, 5, 6},
			slice2: []int{7, 8, 9, 10, 11, 12},
			want:   []int{7, 8, 9, 10, 11, 12},
		},
		{
			name:   "repeated values",
			slice1: []int{3, 3, 3, 3, 3, 3, 3},
			slice2: []int{1, 1, 1, 1, 2, 2, 2, 1},
			want:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DiffRight(tt.slice1, tt.slice2)
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

func BenchmarkDiffRight(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{5, 6, 7, 8}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DiffRight(slice1, slice2)
	}
}

func BenchmarkDiffRight_MoreValues_Random(b *testing.B) {
	var slice1, slice2 []int
	for i := 0; i < 300; i++ {
		slice1 = append(slice1, i)
	}

	for i := 0; i < 50; i++ {
		slice2 = append(slice2, i)
	}
	rand.Shuffle(len(slice1), func(i, j int) { slice1[i], slice1[j] = slice1[j], slice1[i] })
	rand.Shuffle(len(slice2), func(i, j int) { slice2[i], slice2[j] = slice2[j], slice2[i] })

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DiffRight(slice2, slice1)
	}
}
