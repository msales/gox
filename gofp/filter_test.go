package gofp_test

import (
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestFilter_Int64(t *testing.T) {
	got := Filter([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(v int64) bool {
		return v%2 == 1
	})

	want := []int64{2, 4, 6, 8, 10}

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

func BenchmarkFilter_Int64(b *testing.B) {
	original := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Filter(original, func(v int64) bool {
			return v%2 == 1
		})
	}
}
