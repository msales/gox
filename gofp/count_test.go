package gofp_test

import (
	"testing"

	. "github.com/msales/gox/gofp"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name       string
		containsFn func() int64
		want       int64
	}{
		{
			name: "count 1",
			containsFn: func() int64 {
				slice := []int{1, 2, 3}
				return Count(slice, func(i int) bool {
					return i == 1
				})
			},
			want: 1,
		},
		{
			name: "count many",
			containsFn: func() int64 {
				slice := []int{1, 2, 3}
				return Count(slice, func(i int) bool {
					return i >= 2
				})
			},
			want: 2,
		},
		{
			name: "count zero",
			containsFn: func() int64 {
				slice := []int{1, 2, 3}
				return Count(slice, func(i int) bool {
					return i < 0
				})
			},
			want: 0,
		},
		{
			name: "count structs",
			containsFn: func() int64 {
				slice := []testStruct{
					{"1", 1},
					{"2", 2},
					{"3", 3},
				}
				return Count(slice, func(ts testStruct) bool {
					return ts.name == "1" || ts.name == "3"
				})
			},
			want: 2,
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
