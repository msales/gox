package gofp_test

import (
	"strconv"
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestMap_Int64ToIn32(t *testing.T) {
	got := Map([]int64{1, 2, 3}, func(val int64) int32 {
		return int32(val)
	})

	want := []int32{1, 2, 3}

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

func TestMap_StringToStruct(t *testing.T) {
	got := Map([]string{"1", "2", "3"}, func(val string) testID {
		id, _ := strconv.ParseInt(val, 10, 64)
		return testID{ID: id}
	})

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

func BenchmarkMap_Int64ToIn32(b *testing.B) {
	original := []int64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Map(original, func(val int64) int32 {
			return int32(val)
		})
	}
}

func BenchmarkMap_StringToStruct(b *testing.B) {
	original := []string{"1", "2", "3"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Map(original, func(val string) testID {
			id, _ := strconv.ParseInt(val, 10, 64)
			return testID{ID: id}
		})
	}
}

func TestMapMultiple_Int64ToIn32(t *testing.T) {
	got := MapMultiple([]int64{1, 2, 3}, func(val int64) []int32 {
		return []int32{
			int32(val),
			int32(val),
		}
	})

	want := []int32{1, 1, 2, 2, 3, 3}

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

func TestMapMultiple_StringToStruct(t *testing.T) {
	got := MapMultiple([]string{"1", "2", "3"}, func(val string) []testID {
		id, _ := strconv.ParseInt(val, 10, 64)
		return []testID{
			{ID: id},
			{ID: id},
		}
	})

	want := []testID{
		{ID: 1},
		{ID: 1},
		{ID: 2},
		{ID: 2},
		{ID: 3},
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

func BenchmarkMapMultiple_Int64ToIn32(b *testing.B) {
	original := []int64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MapMultiple(original, func(val int64) []int32 {
			return []int32{
				int32(val),
				int32(val),
			}
		})
	}
}

func BenchmarkMapMultiple_StringToStruct(b *testing.B) {
	original := []string{"1", "2", "3"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MapMultiple(original, func(val string) []testID {
			id, _ := strconv.ParseInt(val, 10, 64)
			return []testID{
				{ID: id},
				{ID: id},
			}
		})
	}
}

type testID struct {
	ID int64
}
