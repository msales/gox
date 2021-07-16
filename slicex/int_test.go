package slicex_test

import (
	"math/rand"
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestIntToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []interface{}{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
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
			got := IntToInterface(tt.input)

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

func TestIntToString(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []string
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "empty input",
			input: []int{},
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
			got := IntToString(tt.input)

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

func TestIntToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int32
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []int32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
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
			got := IntToInt32(tt.input)

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

func TestIntToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int64
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []int64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
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
			got := IntToInt64(tt.input)

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

func TestIntToFloat32(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []float32
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []float32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
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
			got := IntToFloat32(tt.input)

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

func TestIntToFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []float64
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []float64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
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
			got := IntToFloat64(tt.input)

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

func TestContainsAtLeastOneInt(t *testing.T) {
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
			want:     true,
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
			got := ContainsAtLeastOneInt(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAtLeastOneInt_Single(b *testing.B) {
	haystack := []int{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneInt(haystack, 3)
	}
}

func BenchmarkContainsAtLeastOneInt_Multiple(b *testing.B) {
	haystack := []int{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsAtLeastOneInt(haystack, 2, 3)
	}
}

func TestIntSymDiff(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntSymDiff(tt.slice1, tt.slice2)
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

func BenchmarkIntSymDiff_Both(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{5, 6, 7, 8, 9, 10}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntSymDiff(slice1, slice2)
	}
}

func BenchmarkIntSymDiff_Both_Pessimistic(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntSymDiff(slice1, slice2)
	}
}

func BenchmarkIntSymDiff_Both_MoreValues_Random(b *testing.B) {
	var slice1, slice2 []int
	for i := 0; i < 300; i++ {
		slice1 = append(slice1, i)
	}

	for i := 0; i < 50; i++ {
		slice2 = append(slice2, i)
	}
	rand.Shuffle(len(slice1), func(i, j int) { slice1[i], slice1[j] = slice1[j], slice1[i] })
	rand.Shuffle(len(slice2), func(i, j int) { slice2[i], slice2[j] = slice2[j], slice2[i] })
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntSymDiff(slice1, slice2)
	}
}

func BenchmarkIntSymDiff_Both_Pessimistic_RightShorter(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	slice2 := []int{17, 18, 19, 20}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntSymDiff(slice1, slice2)
	}
}

func TestIntDiffLeft(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntDiffLeft(tt.slice1, tt.slice2)
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

func BenchmarkIntDiffLeft(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{5, 6, 7, 8}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntDiffLeft(slice1, slice2)
	}
}

func BenchmarkIntDiffLeft_MoreValues_Random(b *testing.B) {
	var slice1, slice2 []int
	for i := 0; i < 300; i++ {
		slice1 = append(slice1, i)
	}

	for i := 0; i < 50; i++ {
		slice2 = append(slice2, i)
	}
	rand.Shuffle(len(slice1), func(i, j int) { slice1[i], slice1[j] = slice1[j], slice1[i] })
	rand.Shuffle(len(slice2), func(i, j int) { slice2[i], slice2[j] = slice2[j], slice2[i] })
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntDiffLeft(slice1, slice2)
	}
}

func TestIntDiffRight(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntDiffRight(tt.slice1, tt.slice2)
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

func BenchmarkIntDiffRight(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{5, 6, 7, 8}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntDiffRight(slice1, slice2)
	}
}

func BenchmarkIntDiffRight_MoreValues_Random(b *testing.B) {
	var slice1, slice2 []int
	for i := 0; i < 300; i++ {
		slice1 = append(slice1, i)
	}

	for i := 0; i < 50; i++ {
		slice2 = append(slice2, i)
	}
	rand.Shuffle(len(slice1), func(i, j int) { slice1[i], slice1[j] = slice1[j], slice1[i] })
	rand.Shuffle(len(slice2), func(i, j int) { slice2[i], slice2[j] = slice2[j], slice2[i] })
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IntDiffRight(slice2, slice1)
	}
}
