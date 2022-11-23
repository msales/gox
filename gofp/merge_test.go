package gofp_test

import (
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestMerge_Int64(t *testing.T) {
	got := Merge([]int64{1, 2, 3}, []int64{4, 5, 6})

	want := []int64{1, 2, 3, 4, 5, 6}

	if len(want) != len(got) {
		t.Errorf("Got %+v, want %+v", got, want)
		return
	}

	for k, v := range want {
		if v != got[k] {
			t.Errorf("Got %+v, want %+v", got, want)
		}
	}
}

func TestMerge_Struct(t *testing.T) {
	got := Merge([]testID{{ID: 1}, {ID: 2}}, []testID{{ID: 3}})

	want := []testID{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}

	if len(want) != len(got) {
		t.Errorf("Got %+v, want %+v", got, want)
		return
	}

	for k, v := range want {
		if v != got[k] {
			t.Errorf("Got %+v, want %+v", got, want)
		}
	}
}

func BenchmarkMerge_Int64(b *testing.B) {
	slice1 := []int64{1, 2, 3}
	slice2 := []int64{4, 5, 6}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Merge(slice1, slice2)
	}
}

func BenchmarkMerge_Struct(b *testing.B) {
	slice1 := []testID{{ID: 1}, {ID: 2}, {ID: 3}}
	slice2 := []testID{{ID: 4}, {ID: 5}, {ID: 6}}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Merge(slice1, slice2)
	}
}
