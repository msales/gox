package gofp_test

import (
	"strconv"
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestReindex_Int64ToIn32(t *testing.T) {
	got := Reindex(map[int64]int64{1: 1, 2: 2, 3: 3}, func(key, val int64) (int32, int32) {
		return int32(key), int32(val)
	})

	want := map[int32]int32{1: 1, 2: 2, 3: 3}

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

func TestReindex_StringToStruct(t *testing.T) {
	got := Reindex(map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}, func(key, val string) (string, testID) {
		id, _ := strconv.ParseInt(val, 10, 64)
		return key, testID{ID: id}
	})

	want := map[string]testID{
		"1": {ID: 1},
		"2": {ID: 2},
		"3": {ID: 3},
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

func BenchmarkReindex_Int64ToIn32(b *testing.B) {
	original := map[int64]int64{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Reindex(original, func(key, val int64) (int32, int32) {
			return int32(key), int32(val)
		})
	}
}

func BenchmarkReindex_StringToStruct(b *testing.B) {
	original := map[string]string{
		"1": "1",
		"2": "2",
		"3": "2",
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Reindex(original, func(key, val string) (string, testID) {
			id, _ := strconv.ParseInt(val, 10, 64)
			return key, testID{ID: id}
		})
	}
}
