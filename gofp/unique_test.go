package gofp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/msales/gox/gofp"
)

func TestIsUnique_Int(t *testing.T) {
	type test struct {
		name string
		data []int
		want bool
	}
	tests := []test{
		{
			name: "unique",
			data: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "not unique",
			data: []int{1, 2, 3, 4, 5, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gofp.IsUnique(tt.data)
			assert.Equalf(t, tt.want, got, "Got %+v, want %+v", tt.want, got)
		})
	}
}

func TestIsUnique_String(t *testing.T) {
	type test struct {
		name string
		data []string
		want bool
	}
	tests := []test{
		{
			name: "unique",
			data: []string{"foo", "bar", "baz"},
			want: true,
		},
		{
			name: "not unique",
			data: []string{"foo", "bar", "baz", "foo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gofp.IsUnique(tt.data)
			assert.Equalf(t, tt.want, got, "Got %+v, want %+v", tt.want, got)
		})
	}
}

func TestIsUniqueWithKeyBuilder(t *testing.T) {
	type data struct {
		foo string
		bar []int
	}
	type test struct {
		name string
		d    []data
		want bool
	}
	tests := []test{
		{
			name: "unique",
			d:    []data{{foo: "m"}, {foo: "x", bar: []int{1}}, {foo: "z", bar: []int{8, 9}}},
			want: true,
		},
		{
			name: "not unique ",
			d:    []data{{foo: "m"}, {foo: "x", bar: []int{1}}, {foo: "x", bar: []int{8, 9}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gofp.IsUniqueWithKeyBuilder(tt.d, func(t data) string { return t.foo })
			assert.Equalf(t, tt.want, got, "Got %+v, want %+v", tt.want, got)
		})
	}
}

func BenchmarkIsUnique_Int(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.IsUnique(slice1)
	}
}

func BenchmarkIsUnique_Int_Optimistic(b *testing.B) {
	slice1 := []int{1, 1, 2, 3, 4, 5}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.IsUnique(slice1)
	}
}

func BenchmarkIsUnique_Int_Pessimistic(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 1}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.IsUnique(slice1)
	}
}

func BenchmarkIsUniqueWithKeyBuilder_Int(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.IsUnique(slice1)
	}
}

func BenchmarkIsUniqueWithKeyBuilder_Int_Optimistic(b *testing.B) {
	slice1 := []int{1, 1, 2, 3, 4, 5}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.IsUnique(slice1)
	}
}

func BenchmarkIsUniqueWithKeyBuilder_Int_Pessimistic(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5, 1}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.IsUnique(slice1)
	}
}
