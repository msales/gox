package gofp_test

import (
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		name  string
		init  int64
		slice []int64
		want  int64
	}{
		{
			name:  "reduce multiple elements",
			init:  0,
			slice: []int64{1, 2, 3, 4, 5, 6},
			want:  21,
		},
		{
			name:  "reduce multiple elements with init value",
			init:  1,
			slice: []int64{2137, 666},
			want:  2804,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.slice, func(accum, item int64) int64 { return accum + item }, tt.init)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkReduce(b *testing.B) {
	original := []int64{1, 2, 3, 4, 5, 6}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Reduce(original, func(accum, item int64) int64 { return accum + item }, 0)
	}
}
