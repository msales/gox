package slicex_test

import (
	"reflect"
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name       string
		containsFn func() bool
		want       bool
	}{
		{
			name: "found simple",
			containsFn: func() bool {
				slice := []int{1, 2, 3}
				return Contains(slice, func(i int) bool {
					return slice[i] == 1
				})
			},
			want: true,
		},
		{
			name: "not found simple",
			containsFn: func() bool {
				slice := []int{1, 2, 3}
				return Contains(slice, func(i int) bool {
					return slice[i] == 4
				})
			},
			want: false,
		},
		{
			name: "found struct",
			containsFn: func() bool {
				slice := []testStruct{
					{"1", 1},
					{"2", 2},
					{"3", 3},
				}
				return Contains(slice, func(i int) bool {
					return slice[i].name == "1"
				})
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.containsFn()
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

type containsAtLeastOneTestCase[T comparable] struct {
	name     string
	slice    []T
	contains []T
	want     bool
}

func TestContainsAtLeastOne_GenericFloat64(t *testing.T) {
	tests := []containsAtLeastOneTestCase[float64]{
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
			want:     true,
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

	testContainsAtLeastOne[float64](t, tests)
}

func TestContainsAtLeastOne_GenericInt64(t *testing.T) {
	tests := []containsAtLeastOneTestCase[int64]{
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

	testContainsAtLeastOne[int64](t, tests)
}

func TestContainsAtLeastOne_GenericString(t *testing.T) {
	tests := []containsAtLeastOneTestCase[string]{
		{
			name:     "found 1 element",
			slice:    []string{"1", "2", "3", "4", "5", "6"},
			contains: []string{"3"},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []string{"1", "2", "3", "4", "5", "6"},
			contains: []string{"3", "5", "6"},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []string{"1", "2", "3", "4", "5", "6"},
			contains: []string{"3", "5", "6", "7"},
			want:     true,
		},
		{
			name:     "found no elements",
			slice:    []string{"1", "2", "3", "4", "5", "6"},
			contains: []string{"7", "9"},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []string{"1", "2", "3", "4", "5", "6"},
			contains: []string{"7"},
			want:     false,
		},
	}

	testContainsAtLeastOne[string](t, tests)
}

func testContainsAtLeastOne[T comparable](t *testing.T, tests []containsAtLeastOneTestCase[T]) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsAtLeastOne(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestContains_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Contains should panic when not slice is passed")
		}
	}()

	Contains(1, func(i int) bool {
		return true
	})
}

func BenchmarkContains_Simple(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Contains(haystack, func(i int) bool {
			return haystack[i] == 3
		})
	}
}

func BenchmarkContains_Struct(b *testing.B) {
	haystack := []testStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"6", 5},
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Contains(haystack, func(i int) bool {
			return haystack[i].name == "4"
		})
	}
}

func TestFilterInt_Filtered(t *testing.T) {
	slice := []int{1, 2, 3}
	want := []int{2, 3}
	Filter(&slice, func(i int) bool {
		return slice[i] == 1
	})
	if !reflect.DeepEqual(want, slice) {
		t.Errorf("Got %+v, want %+v", slice, want)
	}
}

func TestFilterInt_AllFiltered(t *testing.T) {
	slice := []int{2, 4, 6, 8}
	want := []int{}
	Filter(&slice, func(i int) bool {
		return slice[i]%2 == 0
	})
	if !reflect.DeepEqual(want, slice) {
		t.Errorf("Got %+v, want %+v", slice, want)
	}
}

func TestFilterInt_NotFiltered(t *testing.T) {
	slice := []int{1, 2, 3}
	want := []int{1, 2, 3}

	Filter(&slice, func(i int) bool {
		return slice[i] == 4
	})
	if !reflect.DeepEqual(want, slice) {
		t.Errorf("Got %+v, want %+v", slice, want)
	}
}

func TestFilterStruct_Filtered(t *testing.T) {
	slice := []testStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
	}
	want := []testStruct{
		{"2", 2},
		{"3", 3},
	}
	Filter(&slice, func(i int) bool {
		return slice[i].name == "1"
	})
	if !reflect.DeepEqual(want, slice) {
		t.Errorf("Got %+v, want %+v", slice, want)
	}
}

func TestFilterStruct_NotFiltered(t *testing.T) {
	slice := []testStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
	}
	want := []testStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
	}
	Filter(&slice, func(i int) bool {
		return slice[i].name == "4"
	})
	if !reflect.DeepEqual(want, slice) {
		t.Errorf("Got %+v, want %+v", slice, want)
	}
}

func TestFilter_Int_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Contains should panic when not pointer to slice is passed")
		}
	}()

	Filter(1, func(i int) bool {
		return true
	})
}

func TestFilter_Slice_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Contains should panic when not pointer to slice is passed")
		}
	}()

	Filter([]int{1, 2, 3}, func(i int) bool {
		return true
	})
}

func BenchmarkFilter_Simple(b *testing.B) {
	haystack := []int32{1, 2, 3, 4, 5, 6}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Filter(&haystack, func(i int) bool {
			return haystack[i] == 3 || haystack[i] == 6
		})
	}
}

func BenchmarkFilter_Struct(b *testing.B) {
	haystack := []testStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"6", 5},
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Filter(&haystack, func(i int) bool {
			return haystack[i].name == "3" || haystack[i].value == 4
		})
	}
}

type testStruct struct {
	name  string
	value int
}
